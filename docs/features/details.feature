# file: ./docs/features/details.feature

Feature: Details

    Scenario: should save details to warehouse
        Given there is a device
            | id | 34 |
        And it is assigned to customer
            | id   | 4                |
            | name | testing customer |
        When it sends details to the gateway
            | name | my device |
            | time | now       |
        Then details should be saved to the warehouse

    Scenario: should not save details with no customer
        Given there is a device
            | id | 56 |
        When it sends details to the gateway
            | name | my device |
            | time | now       |
        Then the details should not be saved to the warehouse

    Scenario: should not save details with customer with no details
        Given there is a device
            | id | 57 |
        And it is assigned to customer
            | id | 5 |
        When it sends details to the gateway
            | name | my device |
            | time | now       |
        Then  the details should not be saved to the warehouse
