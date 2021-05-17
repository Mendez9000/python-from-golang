# foo.py
def foo(*args, **kwargs):
    s = "response from Python args=%s kwds=%s" % (args, kwargs)
    return s