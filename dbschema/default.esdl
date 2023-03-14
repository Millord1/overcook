module default {
    type Dish {
        required property title -> str;
        required property duration -> int16;
        property description -> str;
        property comment -> str;
        property image -> str;
        multi link ingredients -> Ingredient {
            property quantity -> int32;
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