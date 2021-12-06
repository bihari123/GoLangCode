package electionday

import "fmt"

// NewVoteCounter returns a new vote counter with
// a given number of inital votes.
func NewVoteCounter(initialVotes int) *int {
    var voteCount *int
    voteCount=&initialVotes
    return voteCount
	panic("Please implement the NewVoteCounter() function")
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
    voteCount:= 0
    
    if counter != nil{
      voteCount = *counter
    
    }
    return voteCount
	panic("Please implement the VoteCount() function")
}

// IncrementVoteCount increments the value in a vote counter
func IncrementVoteCount(counter *int, increment int) {
    *counter += increment
	return
	panic("Please implement the IncrementVoteCount() function")
}

// NewElectionResult creates a new election result
func NewElectionResult(candidateName string, votes int) *ElectionResult {
    var electionResult *ElectionResult
	electionResult =&ElectionResult{candidateName,votes}

    return electionResult
	panic("Please implement the NewElectionResult() function")
}

// DisplayResult creates a message with the result to be displayed
func DisplayResult(result *ElectionResult) string {
    return fmt.Sprintf("%s (%d)",result.Name, result.Votes)


}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map
func DecrementVotesOfCandidate(results map[string]int, candidate string) {

	results[candidate]-=1


}
