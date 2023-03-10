module default {
    type Dish {
        required property title -> str;
        required property duration -> int64;
        property description -> str;
        property comment -> str;
        property image -> str;
        required multi link ingredients -> Ingredient {
            property quantity -> int64;
            property unity -> str;
        }
        multi link steps -> Step
    }

    type Ingredient {
        required property name -> str;
    }

    type Step {
        required property content -> json;
        property comment -> str
    }
}