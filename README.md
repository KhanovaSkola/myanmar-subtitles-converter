# Fixing encoding of Myanmar subtitles on YouTube

From Ashley from the Myanmar team:

> I’m not sure if font is an issue with any other language but for Myanmar it’s been a long decade nightmare. Long story short, most people in Myanmar use a font called Zawgyi that is not Unicode so in the beginning of our project, most translators typed in Zawgyi to do the subtitles. Now we have potentially hundreds of videos with Zawgyi subtitles that cannot be read if one does not have that particular font installed in their computer. It appears broken like the screen shot below.
> I was planning on eventually converting all the Zawgyi fonts to Unicode but now that we have migrated to YouTube from Amara, I cannot download and upload srt files as I had used to. Any thoughts on the solution for this?

JIRA-ticket: [IC-635](https://khanacademy.atlassian.net/browse/IC-635)
[i18n forum discussion](https://international-forum.khanacademy.org/t/fixing-encoding-of-locked-subtitles-on-youtube/1697)

To detect the Zawgyi encoding, we're usign the [ML model from Google](https://github.com/google/myanmar-tools).

To convert from Zawgyi encoding to Unicode, we're using the [Go version of Rabbit Converter](https://github.com/Rabbit-Converter/Rabbit-Go).

To download/upload subs to/from YouTube, we're using [scripts from Khanova skola](https://github.com/khanovaskola/kstools), imported as a git submodule.

## Installation

Get dependencies
```bash
$ git submodule update --init --recursive
$ go get -u github.com/google/myanmar-tools/clients/go
$ go get -u github.com/Rabbit-Converter/Rabbit-Go
```
Compile codes
```bash
$ go build convert_zawgyi.go
$ go build detect_zawgyi.go
```
## Usage

### Download original subtitles

```bash
$ kstools/download_subs.py myanmar_ytids.dat -l my -d subs_original/
```

### Detect zawgyi encoding

In all cases the detector output was either 0.0000 or 1.0000, or -Inf for two faulty subtitles.

```bash
$ ./detect_zawgyi > zawgyi_score.dat
$ awk '{if ($2 > 0.9) print}' zawgyi_score.dat > zawgyi_ytids.dat
```

### Convert to Unicode

```bash
$ ./convert_zawgyi
```

(new subs will be in folder `subs_converted`)

### Upload unicode subtitles to YouTube

```bash
kstools/sync_subs_file2yt.py -l my -u -d subs_converted/ zawgyi_ytids.dat
```
