terraform {
  required_providers {
    favote = {
      version = "0.0.1"
      source  = "codershangfeng/favote"
    }
  }
}

data "favote_votes" "all" {}

# Returns all favorites votes
output "all_votes" {
  value = data.favote_votes.all.votes
}
