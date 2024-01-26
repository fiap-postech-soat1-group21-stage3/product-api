Feature: Product API

  # Scenario: Creating a new product
  #   Given the following product details
  #     | Name         | Description            | Category    | Price |
  #     | Batata frita | Batata frita temperada | sides       | 10.00 |
  #   When a request is made to create the product
  #   Then the response should have status code 201
  #   And the response body should match the expected product details

  # Scenario: Updating an existing product
  #   Given the following product details
  #     | Name         | Description             | Category    | Price |
  #     | Batata frita | Batata frita com molho  | sides       | 15.00 |
  #   And an existing product with ID "8c2b51bf-7b4c-4a4b-a024-f283576cf190"
  #   And a request is made to update the product
  #   Then the response should have status code 200

  # Scenario: Deleting an existing product
  #   Given an existing product with ID "8c2b51bf-7b4c-4a4b-a024-f283576cf190"
  #   When a request is made to delete the product
  #   Then the response should have status code 20

  # Scenario: Getting a list of products
  #   When a request is made to get the list of products
  #   Then the response should have status code 200
  #   And the response body should contain a list of products with details
  #     | Name         | Description            | Category     | Price |
  #     | Batata frita | Batata frita com molho | sides        | 15.00 |
  #     | Sorvete      | Sorvete de chocolate   | sweets       | 5.00 |
  #   And the response body should have the correct count
