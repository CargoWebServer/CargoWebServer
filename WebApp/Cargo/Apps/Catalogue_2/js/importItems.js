/**
 * Cette fonction permet d'importer les information des filtres. 
 */
function getFiltersMapping(filePath) {
    var mapping0 = {
        "TYPENAME": "CatalogSchema.FiltreRectangulaire",
        "OVMM": "M_id",
        "Sorte": { "M_Type": { "Sorte": "M_valueOf" } },
        "Secteur": { "M_Secteur": { "Secteur": "M_valueOf" } },
        "Couleur": "M_Couleur",
        "Dure de vie": "M_Duré_de_vie",
        "Keywords": "M_keywords",
        "Commentaires": "M_comments",
        "Longueur": { "M_Longueur": { "Longueur": "M_valueOf", "LongueurUnit": { "M_unitOfMeasure": { "LongueurUnit": "M_valueOf" } } } },
        "Largeur": { "M_Largeur": { "Largeur": "M_valueOf", "LargeurUnit": { "M_unitOfMeasure": { "LargeurUnit": "M_valueOf" } } } },
        "Epaisseur": { "M_Épaisseur": { "Epaisseur": "M_valueOf", "EpaisseurUnit": { "M_unitOfMeasure": { "EpaisseurUnit": "M_valueOf" } } } }
    }

    // The validator, if is true the object will be created.
    mapping0.isValid = function (data) {
        // Here each value must have a mapping...
        return data[14] == "Rectangulaire" // The column 14 represente la forme du filtre...
    }

    var mapping1 = {
        "TYPENAME": "CatalogSchema.FiltreRond",
        "OVMM": "M_id",
        "Sorte": { "M_Type": { "Sorte": "M_valueOf" } },
        "Secteur": { "M_Secteur": { "Secteur": "M_valueOf" } },
        "Couleur": "M_Couleur",
        "Dure de vie": "M_Duré_de_vie",
        "Keywords": "M_keywords",
        "Commentaires": "M_comments",
        "Diameter": { "M_Diamètre": { "Diameter": "M_valueOf", "DiameterUnit": { "M_unitOfMeasure": { "DiameterUnit": "M_valueOf" } } } },
        "Epaisseur": { "M_Épaisseur": { "Epaisseur": "M_valueOf", "EpaisseurUnit": { "M_unitOfMeasure": { "EpaisseurUnit": "M_valueOf" } } } }
    }

    // The validator, if is true the object will be created.
    mapping1.isValid = function (data) {
        // Here each value must have a mapping...
        return data[14] == "Rond" // The column 14 represente la forme du filtre...
    }

    // same as mapping 0 but with different validator.
    var mapping2 = Object.assign({}, mapping0); // Deep copy
    mapping2.TYPENAME = "CatalogSchema.FiltreTube"

    // The validator, if is true the object will be created.
    mapping2.isValid = function (data) {
        // Here each value must have a mapping...
        return data[14] == "Rouleau" // The column 14 represente la forme du filtre...
    }

    var mapping3 = {
        "TYPENAME": "CatalogSchema.FiltrePoche",
        "OVMM": "M_id",
        "Sorte": { "M_Type": { "Sorte": "M_valueOf" } },
        "Secteur": { "M_Secteur": { "Secteur": "M_valueOf" } },
        "Couleur": "M_Couleur",
        "Dure de vie": "M_Duré_de_vie",
        "Keywords": "M_keywords",
        "Commentaires": "M_comments",
        "Diameter": { "M_Diamètre": { "Diameter": "M_valueOf", "DiameterUnit": { "M_unitOfMeasure": { "DiameterUnit": "M_valueOf" } } } },
        "Longueur": { "M_Longueur": { "Longueur": "M_valueOf", "LongueurUnit": { "M_unitOfMeasure": { "LongueurUnit": "M_valueOf" } } } },
        "Grade": { "M_Grade": { "Grade": "M_valueOf", "GradeUnit": { "M_unitOfMeasure": { "GradeUnit": "M_valueOf" } } } }
    }

    // The validator, if is true the object will be created.
    mapping3.isValid = function (data) {
        // Here each value must have a mapping...
        return data[14] == "Poche" // The column 14 represente la forme du filtre...
    }

    return [mapping0, mapping1, mapping2, mapping3]
}