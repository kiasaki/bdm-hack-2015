package pool

// PoolConn is a wrapper around Stoppable to modify the the behavior of
// Stoppable's Close() method.
type PoolConn struct {
	Stoppable
	c        *channelPool
	unusable bool
}

// Close() puts the given connects back to the pool instead of closing it.
func (p PoolConn) Close() error {
	if p.unusable {
		if p.Stoppable != nil {
			p.Stoppable.Stop()
			return nil
		}
		return nil
	}
	return p.c.put(p.Stoppable)
}

// MarkUnusable() marks the connection not usable any more, to let the pool close it instead of returning it to pool.
func (p *PoolConn) MarkUnusable() {
	p.unusable = true
}

// newConn wraps a standard Stoppable to a poolConn Stoppable.
func (c *channelPool) wrapConn(conn Stoppable) Stoppable {
	p := &PoolConn{c: c}
	p.Stoppable = conn
	return p
}
