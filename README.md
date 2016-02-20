# SpaceX Stats parser
Small command line utility that parses the webpage spacexstats.com and prints data about upcoming spacex launches + live countdown.

## Data:
All data is owned and provided by spacexstats.com

## Usage:
```bash
    $ cd ./spacex
    $ go install
    $ spacex
```

Obs: make sure your terminal is wide enough, the countdown uses \r to return to the start of the line every seconds to reprint the content. Thus if the terminal isn't wide enough the content will get garbled.

## Output example
```bash
$ spacex
# Next Launch - SES-9
SpaceX's second launch for SES, lofting a 5300kg communications satellite that will provide SES with more coverage over Southeast Asia.

Status: Upcoming
Launch time: 2016-02-24 UTC
Created: 2016-02-08 09:29 UTC
Updated: 2016-02-08 09:29 UTC

Countdown until Launch: 3 days, 23 hours, 14 minutes, 0 seconds (specific time unknown)
```

## License
Copyright (c) 2016 Kristoffer A. Iversen All rights reserved.
Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
