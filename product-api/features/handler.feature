# Feature: Customer API Handler
#   Scenario: Registering routes for Customer API
#     Given a new customer handler
#     When registering routes for API version "/api/v1"
#     Then the registered routes should match the expected routes:
#       | Method | Path                    |
#       | POST   | /api/v1/customer/       |
#       | GET    | /api/v1/customer/:cpf   |
