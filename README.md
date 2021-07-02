# Geliştirilmiş Video Çözüm

Mobile app of Fernus is great but website... it is not, there is few reasons;

- you have to click 2 times while changing videos, which is really annoying when you want to watch lots of questions
- can't change the speed of video, some teachers speak slow or maybe you just want to take a quick look
- doesn't have any accesibility with keyboard
- sometimes website sends file not found error, in response a alert() happens with base64 error message, you have to refresh website and select book again

and a few more... i just want to make something better.

## Site Design

- when path changed site path changes so user don't need to select book every time website opens
- after book selected; subjects, tests and question numbers shown at left and video at right, you can quickly change questions
- left-arrow: previus question, right-arrow: next question
- black and white theme

## Video Player

- you can select between 0.25x to 2x speeds
- autoplay
- J: -5s, K: +5s (same as youtube)

## Technical

- no more base64 error that sometimes happen (mostly on hizverenk)
- using [Ruffle](https://ruffle.rs/) instead of a endpoint that converts swf to jpg, if wasm not supported it fallbacks to server generated png

TODO: PUT WEBSITE SPEED TESTS AND COMPARISONS WITH OLD SITE

## For Developers

Developing: `yarn dev`
Building: `yarn build`
Build Preview: `yarn preview`
