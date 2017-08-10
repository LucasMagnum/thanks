Thanks API
===========

This API was built to implement a positive and simple feedback system using slack.


### TODO LIST
1. How to integrate the slack command
  1.1. Create a slack command
  1.2. Test the slack command
  1.3. Host the API
  1.4. Set allowed companies

2. Slack Commands
  FeedbackCommand
    - Explain how it works
  RankingCommand
    - Explain how it works

3. Development
  How to install
  How to run tests
  How to contribute


### How to use

Once we have the commands installed on our slack, we can run `/thanks` and `/thanks-ranking`:


##### /thanks

The slack command `/thanks` is used to give a positive feedback.

  You've to inform one or more usernames and the reason for your positive feedback:

  ```
   $ /thanks @username for the awesome pair programming
    Good job @username! You've earned +1 point.

   $ /thanks @username1 @username for the code review
    Good job @username1 @username! You've earned +1 point.
  ```


##### /thanks-ranking

This is a kinda feedback gamification, we can use the `/thanks-ranking <today|month|all>` to see the ranking:

  ```
   $ /thanks-ranking today
    Ranking (Today)
    1. @username            1
    2. @username1           1


  $ /thanks-ranking month
    Ranking (Month)
    1. @username           12
    2. @username1          11
  ```

The ranking is available online too, just access http://app-host-address/company-name/


### How it works

  This API can be used to one or many companies at the same time, every information is saved with the company name.

  We can see it at database level:

  ```
    $ desc feedbacks;

    id  from        to          company
    1   @username   @username1  Company AB
  ```


### How to integrate

  1. Clone this project
  2. Host (see Heroku)
  3. Configure slack
  4. Start using


### Development

To run this project, we have to follow the steps bellow:

1. Clone this repository
2. Install all dependencies (see [setup](#1-setup))
3. Run the tests (see [tests](#2-running-tests))


### Table of Contents

  * [1. Setup](#1-setup)
  * [2. Running tests](#2-running-tests)
  * [3. How to test slack command](#3-test-slack-command)
  * [Useful links](#useful-links)

---

### 1. Setup

To properly run this project, you should have Golang > 1.8


  ```bash
  $ go version
  go version go1.8.3
  ```

* Note: this project was build using Golang 1.8.3, but it could work in old versions too.


### 2. Running tests

To run tests, just execute `make tests`:

  ```bash
  $ make tests
  ```


### Useful links

  1. [Install Go](https://golang.org/doc/install)
  2. [Go for beginners (Portuguese Version)](https://medium.com/@lucasmagnum/iniciando-em-go-6a34d200f02c)

