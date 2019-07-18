# Change Log

## 0.2.7 (2019-07-18)

Integrates size package with races.

### Change List

- Replace race size with race size category
- Add unique traits to each race
- Add unique traits to race description

## 0.2.6 (2019-07-18)

Adds basic monsters.

### Change List

- Add basic monsters and a monster API
- Add size package and use it with monsters and animals

## 0.2.5 (2019-07-14)

Fixes a bug in deity simplification.

### Change List

- Fix a bug in generating deities where sometimes a blank simplified deity would be created

## 0.2.4 (2019-07-01)

Adds translation to languages and population of word lists based
on the culture's origin climate.

### Change List

- Update word list based on culture climate
- Add translation functions
- Add capitalize first function

## 0.2.3 (2019-07-01)

Big update to languages this time around. This is foundational work
for real phrasebooks and better name generation.

### Change List

- Added generation of coherent word lists to languages
- Added a "Rosetta stone" phrase to language output
- Added basic verb conjugation support

## 0.2.2 (2019-07-01)

Fixed two crash bugs.

### Change List

- Fixed a crash when holy items weren't being calculated correctly
- Fixed a crash when no fish were generated for a climate

## 0.2.1 (2019-06-30)

Building styles and clothing styles are the big winners in this release.
Also quite a few bug fixes made it in.

### Change List

- Clothing styles are much improved, with outfits making more sense
- Building styles got a little more variety
- Refactoring was done for animals and plants to make changes easier in the future
- Bug fixes, primarily in materials

## 0.2.0 (2019-06-25)

This version moves to using manual semantic versioning. The previous
scheme used 0.1.X versioning, where X was the CircleCI build number.

The version number 0.2.0 was chosen to differentiate it from that
process.

### Change List

- Tag the current version of the World API as v0.2.0
- Update the CI/CD configuration for CircleCI to reflect the new scheme
- Add the git commit hash as an additional Docker image tag