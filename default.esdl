type Dish {
  required property title -> str
  required property duration -> int64
  required property description -> str
  required property Comment -> str
  required multi link ingredients -> Ingredient {
    required property quantity -> int64
    required property unity -> str
  }
}

type Ingredient {
  required property name -> str
}

type Step {
  required property content -> json
}