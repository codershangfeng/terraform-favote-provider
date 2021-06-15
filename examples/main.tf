terraform {
  required_providers {
    favote = {
      version = "0.0.1"
      source  = "codershangfeng/favote"
    }
  }
}

data "favote_vote" "vote_1" {
  vid = "1"
}

# Returns all favorites votes
output "vote_1" {
  value = data.favote_vote.vote_1
}
