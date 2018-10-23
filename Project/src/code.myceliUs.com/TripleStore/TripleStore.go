package TripleStore

import (
	"encoding/binary"
	"encoding/json"
	"errors"

	//	"log"

	"code.myceliUs.com/Utility"
	"github.com/syndtr/goleveldb/leveldb"
)

// Convertion function
func uint64ToBytes(v uint64) []byte {
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, v)
	return bs
}

func bytesToUint64(v []byte) uint64 {
	return binary.BigEndian.Uint64(v)
}

type Store struct {
	// Where The store will be create.
	path string

	// The name of the store.
	name string

	// Here I will keep the
	dictionary *leveldb.DB

	// Now the three indexation map...
	subjects   *leveldb.DB
	predicates *leveldb.DB
	objects    *leveldb.DB
}

/**
 * Append a value to the dictionary and return it index.
 */
func (store *Store) appendIndexValue(value []byte) uint64 {
	index, err := store.getValueIndex(value)
	if err != nil {
		var lastIndex uint64
		v, err := store.dictionary.Get([]byte("last_index"), nil)

		if err == nil {
			lastIndex = bytesToUint64(v) + 1
		}

		bs := uint64ToBytes(lastIndex)

		// Set the last index to increment it in next insertion.
		store.dictionary.Put([]byte("last_index"), bs, nil)

		// Store the value as key to retreive the index.
		store.dictionary.Put(value, bs, nil)

		// Store the index as key to retreive it value.
		store.dictionary.Put(bs, value, nil)

		return lastIndex
	}
	return index
}

/**
 * Return the index value from the dictionary.
 */
func (store *Store) getIndexValue(index uint64) (string, error) {
	data, err := store.dictionary.Get(uint64ToBytes(index), nil)
	if err != nil {
		return "", err
	}

	return Utility.ToString(data), nil
}

func (store *Store) getValueIndex(value []byte) (uint64, error) {
	v, err := store.dictionary.Get(value, nil)
	if err != nil {
		return uint64(0), err
	}
	return bytesToUint64(v), nil
}

func NewStore(path string, name string) *Store {
	store := new(Store)
	store.name = name
	store.path = path

	// Open the dictionary or create it...
	store.dictionary, _ = leveldb.OpenFile(path+"/"+name, nil)

	// Now the three indexation store.
	store.subjects, _ = leveldb.OpenFile(path+"/"+name+"/s", nil)
	store.predicates, _ = leveldb.OpenFile(path+"/"+name+"/p", nil)
	store.objects, _ = leveldb.OpenFile(path+"/"+name+"/o", nil)

	return store
}

// Index the subject.
func (store *Store) appendIndex(v0 uint64, v1 uint64, v2 uint64, db *leveldb.DB) {
	// Here I will the list of
	values, err := db.Get(uint64ToBytes(v0), nil)
	if err == nil {
		// Here I must append the values only if there not already exist...
		for i := 0; i <= len(values)-16; i += 16 {
			v1_ := bytesToUint64(values[i : i+8])
			v2_ := bytesToUint64(values[i+8 : i+16])
			if v1_ == v1 && v2_ == v2 {
				return // already in the list.
			}
		}
	}

	// No subject already exist for that triple.
	values = append(values, uint64ToBytes(v1)...)
	values = append(values, uint64ToBytes(v2)...)

	// Append the newly created value.
	db.Put(uint64ToBytes(v0), values, nil)
}

// Append triple in the datastore.
func (store *Store) AppendTriple(subject string, predicate string, object interface{}) {
	//log.Println("----> append triple: ", subject, predicate, object)
	// Append the value into the dictionary if there not already exist
	s := store.appendIndexValue([]byte(subject))
	p := store.appendIndexValue([]byte(predicate))

	// in the case of object I will set a uuid of the value
	o_, _ := json.Marshal(object)
	o := store.appendIndexValue(o_)

	// Test if the information already exist...
	spo := Utility.ToString(s) + ":" + Utility.ToString(p) + ":" + Utility.ToString(o)
	exist, _ := store.dictionary.Has([]byte(spo), nil)

	if !exist {
		// if not index it!
		store.dictionary.Put([]byte(spo), nil, nil)

		// Append the subjects index.
		store.appendIndex(s, p, o, store.subjects)

		// Append the predicate index.
		store.appendIndex(p, s, o, store.predicates)

		// Append the object index.
		store.appendIndex(o, s, p, store.objects)
	}

}

func (store *Store) removeIndex(v0 uint64, v1 uint64, v2 uint64, db *leveldb.DB) {

	// I will remove it from existing values.
	values, err := db.Get(uint64ToBytes(v0), nil)
	values_ := make([]byte, 0)
	if err == nil {
		// Here I must append the values only if there not already exist...
		for i := 0; i <= len(values)-16; i += 16 {
			v1_ := bytesToUint64(values[i : i+8])
			v2_ := bytesToUint64(values[i+8 : i+16])
			if v1_ != v1 && v2_ != v2 {
				values_ = append(values_, values[i:i+8]...)
				values_ = append(values_, values[i+8:i+16]...)
			}
		}
	}

	// Append the newly created value.
	if len(values_) > 0 {
		db.Put(uint64ToBytes(v0), values_, nil)
	} else {
		// delete v0 from dictionary
		v, _ := store.getIndexValue(v0)

		// remove it from dictionary
		store.dictionary.Delete([]byte(v), nil)
		store.dictionary.Delete(uint64ToBytes(v0), nil)
		db.Delete(uint64ToBytes(v0), nil)
	}
}

// remove triple
func (store *Store) RemoveTriple(subject string, predicate string, object interface{}) {

	var s, p, o uint64
	var err error
	s, err = store.getValueIndex([]byte(subject))
	if err != nil {
		return
	}
	store.removeIndex(s, p, o, store.subjects)

	p, err = store.getValueIndex([]byte(predicate))
	if err != nil {
		return
	}
	store.removeIndex(p, s, o, store.predicates)
	// Here I will made use of json to serialyse the information and save it in
	// level db as a byte array. So no coversion will be made accidentaly here.
	o_, _ := json.Marshal(object)
	o, err = store.getValueIndex(o_)

	if err != nil {
		return
	}
	store.removeIndex(o, s, p, store.objects)

	// remove the triple itself.
	spo := Utility.ToString(s) + ":" + Utility.ToString(p) + ":" + Utility.ToString(o)
	store.dictionary.Delete([]byte(spo), nil)
	//log.Println("---->remove triple: ", subject, predicate, Utility.ToString(object))

}

// Find triple in the datastore.
func (store *Store) FindTriples(subject string, predicate string, object interface{}) ([][]interface{}, error) {
	results := make([][]interface{}, 0)
	//log.Println("find ", subject, predicate, Utility.ToString(object))
	// If the subject is know
	if subject != "?" {
		s, err := store.dictionary.Get([]byte(subject), nil)
		if err != nil {
			return results, err
		}
		// So here I will search in the predicate store.
		values, err_ := store.subjects.Get(s, nil)
		if err_ != nil {
			return results, err_
		}

		// prdicate can be specified
		hasPredicate := predicate != "?"
		if hasPredicate {
			_, err := store.dictionary.Get([]byte(predicate), nil)
			if err != nil {
				return results, err
			}
		}

		// Object can be specified
		hasObject := object != "?"
		if hasObject {
			o, _ := json.Marshal(object)
			_, err := store.dictionary.Get(o, nil)
			if err != nil {
				return results, err
			}
		}

		for i := 0; i <= len(values)-16; i += 16 {
			p := bytesToUint64(values[i : i+8])
			o := bytesToUint64(values[i+8 : i+16])

			// So here I will append the values.
			p_, err_p := store.getIndexValue(p)
			o_, err_o := store.getIndexValue(o)

			var o__ interface{}
			json.Unmarshal([]byte(o_), &o__)

			if err_p == nil && err_o == nil {
				if !hasObject && !hasPredicate {
					results = append(results, []interface{}{subject, p_, o__})
				} else if hasObject && o__ == object {
					results = append(results, []interface{}{subject, p_, o__})
				} else if hasPredicate && p_ == predicate {
					results = append(results, []interface{}{subject, p_, o__})
				}
			}
		}

	} else if predicate != "?" {
		//log.Println("---> ", subject, predicate, object)
		p, err := store.dictionary.Get([]byte(predicate), nil)
		if err != nil {
			return results, err
		}
		// So here I will search in the predicate store.
		values, err_ := store.predicates.Get(p, nil)
		if err_ != nil {
			return results, err_
		}

		// Subject can be specified
		hasSubject := subject != "?"
		if hasSubject {
			_, err := store.dictionary.Get([]byte(subject), nil)
			if err != nil {
				return results, err
			}
		}

		// Object can be specified
		hasObject := object != "?"
		if hasObject {
			o, _ := json.Marshal(object)
			_, err := store.dictionary.Get(o, nil)
			if err != nil {
				return results, err
			}
		}

		// So here I will get the results.
		for i := 0; i <= len(values)-16; i += 16 {
			s := bytesToUint64(values[i : i+8])
			o := bytesToUint64(values[i+8 : i+16])

			// So here I will append the values.
			s_, err_s := store.getIndexValue(s)
			o_, err_o := store.getIndexValue(o)

			var o__ interface{}
			json.Unmarshal([]byte(o_), &o__)

			if err_s == nil && err_o == nil {
				if !hasSubject && !hasObject {
					results = append(results, []interface{}{s_, predicate, o__})
				} else if hasSubject && subject == s_ {
					results = append(results, []interface{}{s_, predicate, o__})
				} else if hasObject && object == o__ {
					results = append(results, []interface{}{s_, predicate, o__})
				}
			}
		}

	} else if object != "?" {
		o_, _ := json.Marshal(object)
		o, err := store.dictionary.Get(o_, nil)
		if err != nil {
			return results, err
		}

		// So here I will search in the predicate store.
		values, err_ := store.objects.Get(o, nil)
		if err_ != nil {
			return results, err_
		}

		// Subject can be specified
		hasSubject := subject != "?"
		if hasSubject {
			_, err := store.dictionary.Get([]byte(subject), nil)
			if err != nil {
				return results, err
			}
		}

		// prdicate can be specified
		hasPredicate := predicate != "?"
		if hasPredicate {
			_, err := store.dictionary.Get([]byte(predicate), nil)
			if err != nil {
				return results, err
			}
		}

		// So here I will get the results.
		for i := 0; i <= len(values)-16; i += 16 {
			s := bytesToUint64(values[i : i+8])
			p := bytesToUint64(values[i+8 : i+16])

			// So here I will append the values.
			s_, err_s := store.getIndexValue(s)
			p_, err_p := store.getIndexValue(p)

			if err_p == nil && err_s != nil {
				if !hasSubject && !hasPredicate {
					results = append(results, []interface{}{s_, p_, object})
				} else if hasSubject && s_ == subject {
					results = append(results, []interface{}{s_, p_, object})
				} else if hasPredicate && p_ == predicate {
					results = append(results, []interface{}{s_, p_, object})
				}
			}
		}
	}

	if len(results) == 0 {
		err := errors.New("no reuslts found for " + subject + ", " + predicate + " ," + Utility.ToString(object))
		return results, err
	}

	return results, nil
}
