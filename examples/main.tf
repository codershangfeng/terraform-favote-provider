terraform {
  required_providers {
    favote = {
      version = "0.0.1"
      source  = "codershangfeng/favote"
    }
  }
}

data "favote_vote" "vote_1" {
  topic = "What's your favorite sports?"
}

# Returns all favorites votes
output "vote_1" {
  value = data.favote_vote.vote_1
}

resource "favote_vote" "name" {
  topic   = "What's your favorite sports?"
  options = ["Football", "Basketball", "Tennis"]
}