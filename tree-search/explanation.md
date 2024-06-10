The CheckDuplicateIDs implements a Breadth-First Search algorithm to traverse each level of the tree
before moving on to the next level. This approach was chosen as the goal is to identify a duplicate
as shallowly as possible as quickly as possible. Were we to traverse the tree Depth-First, we *might*
find a duplicate faster, but if a branch is particularly long or the duplicate is in a different child 
node, it may take longer. Further, in order to be sure we were getting truly the shallowest duplicate,
we would still need to traverse additional nodes following a successful duplicate indentification. 

Technically the time complexity of this algorithm is O(n) where n is the number of nodes. In practice however, 
if we expect a duplicate to appear, it will be less, as we will eliminate entire branches before traversal. 
Space complexity is relative to the width of the tree O(w), where w is the number of nodes at the widest 
point in the tree. This is a drawback here as this has the potential to be larger thanthe O(h) height-based 
space complexities with Depth-First approaches, depending on the shape of the tree. However, because we expect 
to find the duplicate prior to complete traversal, we would rarely see the maximum complexities.

Alternatives:

- With more contextual information about the incoming data tree, it may be adventageous to adjust the
algorithm to target certain branches first in a Depth-first pattern if we project the duplicate is more 
likely to be present there.

Improvements: 

- To improve space complexity, we can likely find a way to reference values in the queue and their position
rather than dequeuing them. This isn't immediately obvious to me as my experience in Go is not extensive.

- To make the algorithm more extendable and applicable to real-world use-cases it needs to be able to accept 
a true n-ary tree and performantly handle the queueing of n number of child nodes.