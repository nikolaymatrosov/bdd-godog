Feature: simple

  Scenario Outline: demo second word
    Given the string "<value>"
    When I run echo "<value>"
    Then the output should contain "<value>"
    And clear all

    Examples:
      | value |
      | hello |
      | world |

  Scenario: failing
    Given the string "value"
    When fail step
    Then the output should contain "value"