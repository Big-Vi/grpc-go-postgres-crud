table "orders" {
  schema = schema.public
  column "id" {
    null = false
    type = serial
  }
  column "customer_id" {
    null = false
    type = character_varying(100)
  }
  column "product_name" {
    null = false
    type = character_varying(100)
  }
  column "price" {
    null = false
    type = integer
  }
  column "quantity" {
    null = false
    type = integer
  }
  column "created_at" {
    null = true
    type = timestamp
  }
  column "updated_at" {
    null = true
    type = timestamp
  }
}
schema "public" {
}
