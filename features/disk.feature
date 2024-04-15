Feature: create disk
  As a user
  I want to create a disk

  Scenario: create disk
      Given an image with family "ubuntu-2204-lts" in the "standard-images" folder
      When I create a disk from it with name "my-disk" in the "b1gbpo1c8qkicn81mfok" folder
      Then I should see the disk in the folder
      Then I want to delete the disk