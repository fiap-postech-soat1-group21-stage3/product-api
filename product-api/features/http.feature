Feature: Product API

  Scenario: Creating a new product
    Given the following product details
      | Name         | Description            | Category    | Price |
      | Batata frita | Batata frita temperada | sides       | 10.00 |
    When a request is made to create the product
    Then the response should have status code 201
    And the response body should match the expected product details

  Scenario: Getting a list of products
    When a request is made to get the list of products
    Then the response should have status code 200
    And the response body should contain a list of products with details
      | Name         | Description            | Category     | Price |
      | Batata frita | Batata frita temperada | sides        | 15.00 |

