%module(directors="1") xapian

%{
/* go.i: SWIG interface file for the Go bindings
 *
 * Copyright (c) 2005,2006,2008,2009,2011,2012,2018 Olly Betts
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License as
 * published by the Free Software Foundation; either version 2 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA  02110-1301
 * USA
 */
%}

#define XAPIAN_SWIG_DIRECTORS

%rename(Apply) operator();
%ignore operator&;
%ignore operator|;
%ignore operator^;
%ignore operator*;
%ignore operator/;

%ignore Xapian::Compactor::resolve_duplicate_metadata(std::string const &key, size_t num_tags, std::string const tags[]);


%include  xapian-head.i

namespace Xapian{
	class MSet {};
	class ESet {};
	class MSetIterator{};
	class ESetIterator{};
};

//%include ../generic/except.i
%include  xapian-headers.i

%extend Xapian::MSet {
	int size() const {
		return self->size();
	}
}

%extend Xapian::MSetIterator {
	Xapian::Document get_document() const{
		return self->get_document();
	}	
}