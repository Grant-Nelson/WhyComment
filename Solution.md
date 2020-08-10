# Solution

Honestly, how long did it take for you to figure out what this does?
Or did you never figure it out?

Read the comments below and see if you were right.

## The Removed Comments

```Go
// Frequency is a probabilistic algorithm which determines the set of values
// which have occurred the most number of times in a stream in linear time.
// At any time during streaming, the current most frequent values can be gotten.
//
// This is useful for counting large amounts of events occurring in a system
// without requiring a massive amount of time or memory. For example, keeping
// track of which links or buttons are clicked on the most by customers.
//
// This is known as the heavy hitters algorithm:
// See https://en.wikipedia.org/wiki/Streaming_algorithm#Frequent_elements
// And https://en.wikipedia.org/wiki/Count%E2%80%93min_sketch
type Frequency struct {...}

// New creates a new frequency counter. The given size is the maximum amount
// of values to return (the k in top-k values). The given hash table size and hash
// function determines how often value collisions are allowed at the cost of memory.
func New(size, hashTableSize int, hash HashFunction) *Frequency {...}

// Add inserts a value from the event stream to be counted.
func (f *Frequency) Add(data string) {...}

// Results returns the top-k values which have been seen most frequently.
// The results may have less than k values and the results are probabilistic.
// The results may be gotten at anytime while values are streaming.
func (f *Frequency) Results() []string {...}
```

## Comments are important

The point of comments isn't to explain what the code does internally,
normally that isn't useful.

The point of comments is to explain why you would want to use the code.
They help you remember and others learn why the code was written.

After reading the comments, you know what this does and what it is used for.
You also have links to additional information if you want to learn more.

Now ask yourself, did it honestly take longer to read the comments
than to figure it out based on the code?

Finally, I hear all the time that comments get out-of-date. Comments are part
of the code and should be updated. Updates should not be so large that they change
the meaning of why the code should be used anyway. If updating a method makes
the comment out-of-date, you are either abusing a method or the comment was poorly written.

*Note: The Heavy Hitters problem has been used by many companies as an interview question.
If I used this during an interview, I would ask, "What could be done to improve this code?"
The answer I'm looking for isn't example usages or unit-tests (both good) but what I'm looking
for is someone who realizes their code is used by other people and other people need
comments to understand why you wrote that code.*
