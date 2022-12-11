## Star Wars API with Huma

I wanted to play around with some idea around mixing hypermedia with OpenAPI.
So, I wanted to do it with non-work related content, so I decided to implement
the [Star Wars API](https://swapi.dev) but leverage
[`huma`](https://github.com/danielgtaylor/huma).

This is intended as a playground to play around with ideas, so don't take
anything here very seriously.

### Things of Note

As I already mentioned, I'm using Huma. This handles the OpenAPI generation
stuff and also automatically generates schemas for payloads and the URLs they
can be found at. So that covers part of what I'd like. But another part is
going to be adding hyperlinking data to payloads and for that I'm using my toy
project [`go-claxon`](https://github.com/mtiller/go-claxon) which is still very
much in a nascent stage.

I'm also leveraging the [`do`](https://github.com/samber/do) package to handle
dependency injection. This is very useful for web servers and this repository
shows an example of how I've done that using middleware (I may even publish that
middleware separately since it seems quite useful). Along the way, I'm also
using the `do`s sibling package [`lo`](https://github.com/samber/lo) for all
kinds of generics related functionality.

I've included an [`air`](https://github.com/cosmtrek/air) configuration as well.
Using `air` allows you to automatically recompile the server while developing.
Just run `air` on the command line (once `air` is installed) and it takes care
of rebuilding automatically.
