# Big Data Week Montreal 2015 Hack

Hacking on [ghtorrent](http://ghtorrent.org)'s GitHub events data.

# The big plan

In the optimal case (which, now I know means a 1 week hackathon not 24 hours) I
would of tackled those 5 steps in order to gain insigt on **what makes successful
open source projects successful**, **why are some efforts disproportionately more
active and used than others**.

## Step 1: Acquire dataset subset

The full GitHub historiacal event dataset that started being collected in 2012
is 6.5TB in MongoDB or about 600 000 000 documents.

I downloaded the last 2 months incremental dump for a total of about 15Gb.

# Step 2: Injest the raw data efficiently

Next step is parsing and shoving in a message queue the **BSON** formated file's
contents fast enough to allow for rapid iteration.

**BSON** stands for binary-json and is pretty similar to it's insiration except
is has better traversability and is bytes at rest not chars/strings.
