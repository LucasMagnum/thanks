Thanks API
===========

This API was built to implement a positive and simple feedback system using slack.


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

When we type in Slack a command like this:

  (slack) > @lucas.magnum: /thanks @lucas.magnum thanks for building this project

I used my user to write call the `/thanks` command with the text `@lucas.magnum thanks for building this project`, then slack will send a POST request to the configured url similar to:

    POST / {
        "team_domain": "company-1",
        "user_id": "U165XVKJS",
        "user_name": "lucas.magnum",
        "command": "/thanks",
        "text": "<@U165XVKJS|lucas.magnum> thanks for building this project",
        // Other fields ...
    }

This request will be handler by our `handleFeedbackCommand` and we will return the command handler response to slack.


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

### 3. How to test slack command

We can test the command before deploying it, we have some tools to help us =)



### Useful links

  1. [Install Go](https://golang.org/doc/install)
  2. [Go for beginners (Portuguese Version)](https://medium.com/@lucasmagnum/iniciando-em-go-6a34d200f02c)

