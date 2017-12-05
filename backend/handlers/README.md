# Handlers

The http handlers should be kept as small as possible: ideally they should only
extract data from the incoming request, and write output to the response.

Perhaps, to make endpoints (see endpoints package) more sensible, the handlers
should also resolve the token coming from the frontend.
