# plot-uor-skillgain
Explain skillgains on UORenaissance.com shard, by inspecting open source code.

The UOR server is based on RunUO, although with modifications. To produce this
I analyzed the current HEAD revision of the RunUO code, factoring in the
modifications I know about. This is bound to be off in some respects, but it's
at least a baseline from which we can compare empirical data to infer other
modifications. It's also possible the shard maintainers might confirm/deny
differences between UOR and stock RunUO *hint hint* :)

The RunUO source code itself is [here](https://github.com/runuo/runuo).

# Usage
1. Clone this repository to a system with a bash-compatible shell, like so:
   `git clone git@github.com:sayotte/plot-uor-skillgain.git`.
   1. You could also download a zip of the source from [here](https://github.com/sayotte/plot-uor-skillgain/archive/refs/heads/main.zip)
2. Install golang; the command `go version` on your shell should print 
   something useful. 
3. Run `make shell`
4. Run either `make install` or just `go run main.go`
5. ???
6. Profit.

# Donations
Appreciative users are encouraged to praise me on https://uorforums.com, but
this project requires neither pixels nor cash support.
