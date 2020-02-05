# file: ./docs/features/assignments.feature

Feature: Assignments

    Scenario: synchronizer should update customer in kafka if changed in the database
        Given there is a customer
            | id   | 41               |
            | name | testing customer |
        And the customer is changed in the database
            | id   | 41                       |
            | name | testing customer changed |
        Then the synchronizer should update the customer in kafka

    Scenario: synchronizer should remove customer from kafka if deleted in the database
        Given there is a customer
            | id   | 61               |
            | name | testing customer |
        And the customer is removed from the database
        Then the synchronizer should remove the customer from kafka

    Scenario: synchronizer add customer to kafka if added in the database
        Given there is a customer added in the database
            | id   | 71               |
            | name | testing customer |
        Then the synchronizer should update the customer in kafka
