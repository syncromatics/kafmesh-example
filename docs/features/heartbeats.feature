# file: ./docs/features/heartbeats.feature

Feature: Heartbeats

    Scenario: should save heartbeat to warehouse
        Given there is a device
            | id | 24 |
        And it is assigned to customer
            | id   | 14               |
            | name | testing customer |
        When it sends a heartbeat to the gateway
            | time      | now  |
            | isHealthy | true |
        Then the heartbeat should be saved to the warehouse

    Scenario: should not save heartbeat with no customer
        Given there is a device
            | id | 25 |
        When it sends a heartbeat to the gateway
            | time      | now   |
            | isHealthy | false |
        Then the heartbeat should not be saved to the warehouse


    Scenario: should not save heartbeat with customer with no details
        Given there is a device
            | id | 56 |
        And it is assigned to customer
            | id | 15 |
        When it sends a heartbeat to the gateway
            | time      | now  |
            | isHealthy | true |
        Then the heartbeat should not be saved to the warehouse
