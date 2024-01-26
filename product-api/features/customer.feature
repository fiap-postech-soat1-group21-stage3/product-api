# Feature: Customer Operations

#   Background:
#     Given a customer with the following details:
#     | ID  | Name | Email            | CPF          |
#     |     | João | joao@email.com   | 12345678912  |

#   Scenario: Create a customer successfully
#     When the customer is created
#     Then the created customer should have the following details:
#       | ID                                    | Name | Email            | CPF         |
#       | b4dacf92-7000-4523-9fab-166212acc92d  | João | joao@email.com   | 12345678912 |

#   Scenario: Fail to create a customer
#     Given the customer creation fails with an error
#     When the customer is created
#     Then an error should be returned
