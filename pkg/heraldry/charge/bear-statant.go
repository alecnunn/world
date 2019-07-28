package charge

import svg "github.com/ajstarks/svgo"

var BearStatant = Charge{
	Identifier: "bear-statant",
	Name:       "bear statant",
	Noun:       "bear",
	NounPlural: "bears",
	Descriptor: "statant",
	SingleOnly: false,
	Tags: []string{
		"aggressive",
		"animal",
		"bear",
		"strong",
	},
	Render: func(canvas *svg.SVG, hexCode string, lineColor string, canvasWidth int, canvasHeight int, number int) {
		pathData0 := "M101.4,246.6c-2.6,2.1-4.9,4.1-7.6,3c-2.3-1,1-2.9-0.1-4.4c-1.3-0.3-1.5,1.3-2.4,1.6c-1.9,0.5-3.8,2.2-5.9,0.9    c-2.7-1.7,1.5-2.8,0.2-5.4c-3,5.6-9.5,2.8-13.2,6.6c-0.9,1-2.7,1.5-3.5,0.7c-1.5-1.6-0.4-3.8-0.4-5.9c-1.8,0.6-3.7,0.7-5.3,2    c-1.3,1.1-3.2,1.2-4.2,0.1c-0.8-1-0.2-2.7,1.1-3.7c2.3-1.9,3.4-5,5.4-6.7c3.4-2.8,8.1-3.8,12.5-3.4c5.6,0.6,9.5-1.3,12.1-6.1    c0.5-0.8,1.2-1.8,1.2-2.6c-0.6-5.7,1.6-11.1,1.8-16.6c0.3-7.8,1.8-15.5,2.9-23.2c0.1-0.7,0.8-1.7,0.4-2.3    c-5.1-6.7-6-15.3-9.6-22.6c-4.2-8.4-11.2-13.5-20.8-14.7c-3.4-0.4-6.9-0.9-10.2-1.7c-6.1-1.5-10,3.2-15,5    c-2.2,0.8-4.2,2.2-6.4,2.7c-2.6,0.6-3.7,1.8-3.5,4.3c0.1,2.1-1.2,3.4-3,3.4c-1.8,0-2.8-1.3-3-3.4c-0.3-4.2-3.4-9.5-6.9-12.4    c-2.5-2.1-1.8-3.8,0.2-6c6.8-7.8,14.7-14.6,20.2-23.5c2.5-4,4.9-8.1,8.6-11.1c1.4-1.1,0.9-2.5,1-3.9c0.2-3.3-0.4-6.7,0.8-10    c1.6-4.2,4-7.5,8.3-9.2c3.2-1.2,6.2-2.8,9.4-4.2c3.6-1.6,4.5-1.1,4.5,2.7c-0.1,3.5,1.7,6.6,2,11c4.7-7.7,12.3-8.4,19.4-10    c4.3-1,5.1-0.2,5.1,4.2c0,0.9-0.2,1.7,0,2.6c0.6,2.3,0.2,4.3,4.6,3.7c12.4-1.6,25-2.8,37.5-1.9c3.4,0.2,6.8,1.2,10.1,2.2    c10.7,3.4,21.7,5,32.8,6.4c9.8,1.3,19.1,0.2,28.6-1.7c6.8-1.3,13.6,1.2,20.3,1.9c9.5,0.9,17.9,5.2,26.5,8.4    c10.3,3.8,17.3,11.8,21.3,22.1c1.4,3.7,2.5,7.5,4,11.2c2.7,6.5,1.4,13,0.6,19.6c-0.2,1.7-0.7,3.1,0.6,4.7c1.3,1.6,2.9,3.8,0.9,5.5    c-1.8,1.5-4.4,2.5-7,1.1c-3.1-1.6-6.2-3.2-7.1-7.7c-0.6,2.2-0.8,3.5-0.1,4.8c1.3,2.4,1.9,4-0.9,6.6c-3,2.9,0.2,6.8,1.6,10.1    c0.8,1.8,5.8,3.9-0.6,4.8c-0.2,0-0.5,0.7-0.4,1.1c1.6,7.5,0.2,15.4,2.3,22.8c1.4,5.1,3.5,10.1,4.9,15.2c1.5,5.5,0.4,10.4-4.1,14.4    c-2.4,2.2-5.5,2.3-8.1,2.5c-8.2,0.5-16.2,3.2-24.5,2.7c-2.5-0.2-3.6,1.8-4.9,2.8c-3.2,2.6-4.2,1.8-4.3-2.2c-1.9,0.9-3.2,2.6-5.2,2    c-0.5-0.1-1.6,0-1.6-0.1c-0.5-4.7-3.4-1.5-5.2-1.5c-3.3,0-2.3-1.8-1.7-3.5c0.2-0.5,1.9-0.9,0.1-1.6c-0.7-0.3-0.9-0.4-1.7,0.3    c-2.2,1.9-4.9,3.5-7.7,4.4c-1.4,0.4-4.5,2.5-5-1.8c-0.2-2.1-1.7-0.2-2.3,0.1c-1.4,0.6-3,0.9-4,0.1c-1-0.8,0.1-2.3,0.1-3.4    c0.2-3.8,2.3-6.3,4.3-9.2c3-4.4,7.9-2.9,11.9-3.9c6.4-1.6,8.5-6.6,4.4-12c-4-5.3-8-10.6-11.6-16.2c-2.9-4.6-3.5-10.1-4.3-15.3    c-0.2-1.2,0.2-2.7-1.2-3.4c-1.4-0.6-1.9,0.9-2.7,1.6c-1.8,1.5-3.1,1.9-5.3,0c-3.3-3-7.5-4.4-12.1-4.3c-11.7,0-23.3,0-35,0.1    c-1.4,0-4.6-1.3-2.7,2.6c0.1,0.2-0.4,0.7-0.5,1.1c-2.2,4.6-2.5,8.2-1.3,13.9c1,4.5,0.4,9.4-1,13.3c-3,8.3-2.7,17.1-5.5,25.3    c-2.2,6.5-5.6,9.7-12.1,11.1c-6.1,1.3-12.4,2.1-17.4,5.6C100.7,253.4,101.7,249.8,101.4,246.6z"
		pathData1 := "M101.6,121.8c-2.2,0.1-3.6,1.9-5.3,2.9c-0.8,0.5-1.2,3.1-2.7,1.3c-1.1-1.3,1-1.9,1.7-2.7    c2.8-3.2,5.6-6.4,8.6-9.5c0.8-0.9,1.5-1.7,0.6-3.5c-3.3,5.6-8.2,9.5-13,13.3c-1.7,1.4-0.4,2.5-0.1,3.7c1,3.8-1.1,7-4.8,6.2    c-2.2-0.5-3.9,0.4-6,0.3c3.3,2.3,6.6,4.2,10.8,2c-3.8,1.7-7.2,3.7-11.4,3.9c-2.7,0.1-5.3,0.3-8.1,1.1c-2.3,0.6-6.9-0.2-8.8-2.8    c-1.5-2,2.1-2,3.3-3.3c-4.3-0.2-8.4-1.9-12.3,0.6c-0.9,0.5-1.9-0.1-2.7-0.7c-1.8-1.3-3.3-0.5-4.8,0.6c-5.5,4.6-12.3,6.5-19,8.6    c-2.3,0.7-4.2-1.5-3.9-4c0.1-0.4,0.2-1,0-1.3c-2.8-4.2,0.7-6.2,3.1-8.3c7.1-6.1,11-14.6,16.6-21.8c5.7-7.4,12.1-14.3,20.9-18.4    c2.5-1.2,4.4-0.7,6.2,0.6c1.8,1.3,1.3,3.2,0.4,5c-0.5,1.1-0.4,2.1,0.2,3.3c2-0.9,2.9-2,3-4.2c0.3-8.4,2-10.8,10-12.8    c2.6-0.7,5.1-2,8-2c1.8,0,2.5,0.4,2.4,2.4c-0.2,3.2-0.2,6.2-1.5,9.3c-0.5,1.3-2.3,2.5-1.9,4.2c1.5,1,2-1,4.2-1.4    c-3.7,4.3-6.7,7.9-11.4,9.4c-0.8,0.3-1.5,0.6-2.3,0.9c-1.2,0.4-1,1.4-0.7,2c0.4,1,1.1,0.3,1.7,0c1.8-0.9,3.7-1.8,5.7-2.8    c0.1,1.4,0.1,2.6,0.2,3.9c3.1-4.7,7.6-8.7,8.4-14.7c0.2-1.5,0.4-2.5,2.1-2.1c1.8,0.3,2.9-1.2,4.8-1.4c-0.6,2.4-3.6,1.7-4.2,4    c3.5-0.3,5.6-3.2,9-4.4c-1.7,3.8-4.8,6.2-8,9.9c4.2-1.6,5.9-4.5,8.4-6.3c4-2.8,8-4.5,13-4.5c11.7-0.2,22.9,1.2,34.1,4.7    c7.7,2.4,15.9,3.6,24.1,4.1c4.2,0.3,8.4,0.5,12.6,0.7c6.5,0.4,12.7-1.7,19-2.4c2.2-0.3,4.6-0.6,6.8-0.1c9.3,2.2,19,2.8,28,6.7    c7.1,3.1,14.9,5.1,20.5,11c6.2,6.6,9.9,14.4,12.6,22.9c2.1,6.3,2.5,12.5,0.9,18.9c-0.6,2.3-0.5,5,1.4,7.2c0.6,0.7,2.4,1.1,1.2,2.5    c-1.2,1.4-2.9,0.7-4.4,0.6c-1.6-0.1-2-1.5-1.7-2.9c0.5-2.4,1.2-4.7,2-7.8c-3.8,2.2-2.4,5.8-4.3,7.9c-3.3-4.9,0.6-9.3,1.3-14.5    c-3.5,1.8-3.7,4.9-4.7,7.9c-1.5-5.2,1.8-9.5,1.7-14.3c0-2.1,0.2-4.3-0.4-6.3c-0.9-3-2.4-5.8-2.6-9.1c-0.1-1.2-1.1-2.8-2.1-3.4    c-2.8-1.8-4.8-4.3-7.2-6.5c-1.6-1.5-3.3-2.5-5.9-1.8c8.3,6,9.4,13.6,4.9,22.5c3.3-0.9,3.3-0.9,4.8-13.2c3.6,2,4.7,7.2,2.2,10.5    c-0.8,1-2.2,2.3-1.3,3.5c0.7,1.1,2,0.1,2.8-0.7c0.5-0.5,1-0.6,1,0.2c0,0.6,1.1,0.9,0.4,1.7c-0.7,0.8-2.1,1.3-2,2.6    c0.2,2,2.1-0.7,2.4,0.7c0.2,1,0.1,2.3-0.2,3.3c-1.8,5.2-2.2,10.6-2.3,16c-0.1,2.6,2.2,4.5,1.7,7c-3.9,0.5-3-4-6.1-4.5    c6,5.4,1,14.6,8.1,19.8c-4.5,1.1-5.4-2.4-7.3-4.7c-0.2,2.3-0.2,3,0.9,4.4c3.9,4.6,4,9.9,3.6,15.5c-0.3,4.4-3.2,2.8-5.8,2.2    c1,3.8,9.3,1.1,6.2,7.9c-0.5-0.2-1-0.4-1.4-0.7c-0.9-0.9-1.8-1.5-3-0.8c-0.5,0.3-0.3,1.1,0.2,1.4c1.1,0.6,1.5,3.1,3,1.9    c1.6-1.3,2-0.2,2.1,0.7c0.5,2.9,1.9,5.3,3.2,8c-0.5,0-0.7,0.1-0.9,0.1c-1.5-0.1-0.7-3.9-2.8-2.2c-1.6,1.3,1.3,1.8,1.7,3    c1,3,4.1,4.9,4.4,8.3c0.4,5.5-2.6,9.2-8.3,10.1c-7.3,1.2-14.7,2.4-22.3,3.6c2-4.3,6.6-5.1,8.9-8.5c-3.7,1.4-6.8,3.9-9.9,6.2    c-0.8,0.6-1.3,2.2-2.4,1.6c-1.1-0.5-2.3-1.7-2.5-2.7c-0.2-1.5,1.8-1.2,2.7-1.9c0.9-0.8,2.5-0.7,3.1-2.7c-3.6,0.9-6.9,2-9,4.8    c-1.8,2.4-3.5,0.2-5.3,0.5c-1,0.2-1.2-0.7-1-1.4c0.5-1.8,1.9-2.3,3.6-3c2.4-0.9,5.1-1.1,8-3.6c-4.8,1.2-8.6,1.9-12.3,3.3    c-2.1,0.8-2.8,0.2-3.4-1.4c-0.9-2.1,1.5-2.1,2.1-2.2c5.5-0.4,11.1,0.3,16.3-2.6c0.9-0.5,2.6,1.6,4.3,2c1.7,0.5,3.2,0.3,4.9,0.3    c-0.5-2.2-3.4-2.6-3.1-5c0.6-1.7,0-2.6-1.8-2.8c-2.4-3.9-6.4-7-6-12.4c0.2-3.2-3.3-4.5-4.5-7.3c-1-2.3-1.7-4.1-0.3-6.6    c-2.9-0.1-2.8,1.9-3.3,4.1c-2.3-2.8-5.6-4.4-3-8.3c0.2-0.3,0.1-1.2-0.2-1.6c-0.4-0.6-0.9-0.2-1.4,0.1c-1.1,0.6-1.7,3.2-3.4,0.8    c-1.5-2-1.3-4.4-0.3-6.3c1.7-3.3,3.5-2.2,5,1.4c1-3.3-1.2-5.5-0.4-7.9c-2.1,0.2-1.5,1.9-2,2.9c-3.8-1.8-3.5-4.9-2.9-8.4    c-1,1.7-2.8,3.1-1.5,5.3c0.2,0.3,0.4,0.8,0.3,1.2c-0.3,2,1,4.3-1,6c-0.9,0.9-1.7,0-2.2-0.5c-1.8-1.7-1.9-4-2-6.3    c0-2.5,2.3-0.7,3.2-1.8c-2.9-3.1,0-6.5-0.9-10c-0.8,2.2-1.6,4.4-2.2,6.6c-0.2,0.9-0.7,1.9-1.4,1.8c-0.6,0-1-1.1-1.3-2    c-1.6-4.3-0.3-7.8,2.6-11c0.7-0.7,1.1-1.8,1.5-2.6c-1.2-1.2-1.7,1.2-2.6,0.3c-0.5-0.1-0.5-0.3-0.3-0.6c4.2-4.4,8.2-8.8,9.4-15.3    c0.4-2.3,0.4-4.3-0.3-6.5c-0.4-1.2-0.7-2.5-2-2.1c-1.3,0.4-0.8,2.1-0.6,3c1.9,6.5-0.4,11.5-5.1,16c-3.2,3-5.6,6.2-5.4,11.2    c0.1,2.8-3.2,5.6-0.5,9.1c-3.4-2.5-4.6-5.8-6.2-8.8c1.1-1.7,3.5,1.9,4.1-1.4c-3,0.3-5.1-2.1-8-2.7c0,0.6-0.1,0.9-0.1,1.3    c0.2,1.2,3,1.4,1.7,3.1c-1.7,2.1-2.8-0.5-4.2-0.9c-1.5-0.4-2-2.8-4.4-1.3c2.2,2.2,4.2,4.2,7.6,4.7c4.5,0.6,4.1,5.6,6.6,8.2    c-4-1.1-5.9-6.4-9.3-6.5c-4.4-0.1-7.4-3.5-11.7-3.6c3.2,5.2,11.5,2.4,13.3,9.6c-4.6-0.9-7.9-5.8-13.2-4.1c1.9,1.8,5.1,1.4,6.7,3.7    c0.5,0.7,2.2,1,0.8,2.5c-1.1,1.1-1.5,0.8-2.6-0.1c-3.2-2.6-6.8-4.2-11.1-3.9c-5.4,0.4-10.5-0.4-15-3.6c-0.8-0.6-1.2-0.5-1.6,0.2    c-1.3,2.1,1,2.1,2,3.1c-2.5,0.5-4.2-1.1-5.9-2.1c-1.3-0.7-2.5-3.1-3.7-1.6c-1.2,1.6,1.3,2.4,2.8,3.5c-3,1-5.3-0.8-7.7-1.3    c-2.3-0.5-3.3-4.6-6.4-2.2c1,1.9,3,2.3,5.4,3.8c-2.4,0-4.1,0-5.8,0c-2.1,0-3.5-3-6.4-1.1c-0.9,0.6-3.6-1.4-4.8-3.7    c-3-5.6-5.6-11.4-8-18.6c-1.3,3.4-1.5,5.5-0.4,7.4c2.9,5.4,4.6,11.5,9.1,16.1c0.9,1,1.8,2.4,0.9,4.5c-1.6,3.3,0.7,6.8,0.5,10.7    c-1.4-1.2-2-3.6-4.4-1.8c2.6,3.2,5.7,6.3,4.2,11.8c-0.3-0.7-0.4-1.5-0.9-2c-0.7-0.8-0.7-3.1-2.1-2.2c-1.4,0.9-0.4,2.9,0.5,3.8    c3.8,3.7,2.4,7.4,1,11.7c-1.5,4.5-1.5,9.5-2.5,14.2c-1.2,5.7-2.6,11.2-8.3,14.5c-2.5,1.5-4.7,1.5-7.1,0.7    c-2.1-0.7-4.3-0.6-5.5,0.8c-2.6,3-3.8,0.8-5.5-0.7c1.4-2,4.4-1.1,5.6-3.4c-4.5-2.2-7.8,5.3-12.7,1.4c1.8-2.2,4.7-1.9,6.8-3.3    c-4.1-1-7.3,0.1-9.8,3.3c-0.7,0.9-1.6,1.7-2.5,2.3c-0.4,0.3-1.4-0.1-1.9-0.3c-0.4-0.2-1.1-0.8-0.4-1.3c3-2.1,3.8-6.5,9.1-6.1    c8.6,0.6,11.7-4.5,9.4-13.1c-1.9-7.1-3.6-14.2-3.6-21.6c0-3.6,1.3-7.2,0.8-10.9c-0.1-1.1,0.1-2,1-2.6c1.2-0.8,1.7,0.5,2.4,1    c0.7,0.5,1.2,2,2.3,0.8c0.6-0.7,0.4-1.5-0.8-2.3c-1.1-0.6-1.7-2.3-2.6-3.8c3.1-1.4,5.7,5.5,8.2-0.3c-3.5-3.5-7.3-7.7-11.5-11.4    c-3.2-2.9-2.3-6.3-2.5-9.7c-0.1-0.9,0.6-2.3-0.5-2.6c-1.4-0.5-1.6,1-2.2,1.9c-1.8,3-2.5,5.8,0.1,8.8c0.9,1,1.3,2.5,0.8,3.9    c-6-3.9-5.5-12.7,1.3-27.6c-2.8,2.1-5.7,4-6.4,8.2c-0.6,3.5-2.3,6.9-3.8,11.3c-1.3-13.8,7.9-21.9,13.2-31.8    c-3.6,1.3-7.9,6.2-8.4,9.5c-0.1,0.7-0.1,1.5-0.5,2c-3.1,3.2-6.5,6.1-7.3,10.9c-1.1-3.6-0.3-6.1,2.7-8.5c1.5-1.2,5.2-2.4,2.9-6.2    c-0.6-1,2-1.2,2.5-2.7c0.4-1.3,1.9-2.2,3.2-3.6c-5.2,0.7-6.8,5.6-10.7,7.5c-0.5,0.2-0.7,0.8-0.4,1.5c0.7,1.2,1.7-1,2.4,0.3    c-1.6,3.1-5.2,4.8-6.3,8.3c-1.8-2.9-1.7-2.9,0.9-5.1c1.1-1,3.2-1.6,2-3.9c-0.3-0.6-0.7-1.2-1.1-0.9c-2.5,2.1-6.1-1.2-8.5,1.8    c-0.7,0.8-1.6-0.6-2-1.4c4.6-2,9.9-1.5,14.4-3.9c2.1-1.1,2-2.4,1.7-4.4c2,0.9,3.4-0.3,5.1-0.8c6.8-2.2,9.6-9.6,16.4-11.5    c-7.4,1-11.5,7.8-18.2,10c0,0.1-0.1-0.3-0.1-0.6c-0.2-1.3,2.1-3.1,0.3-3.6c-2.4-0.6-1.7,2.2-2.2,3.9c-2.1-1.2-1.6-2.3-0.4-4    c1.6-2.2,4.7-2.6,5.7-5.3l0.2-0.1L101.6,121.8z"
		pathData2 := "M101.8,180c-1-0.4-2.4,0-2.5-2c-0.1-1.2,0.7-4.4-2.7-2.9c-0.2,0.1-1.2-0.8-0.9-1.7c0.3-0.8,1.6-0.7,1.8-0.5    c2.2,2.7,6.3,3,7.7,6.6c0.7,1.8-1.2,2.8-1.4,4.3c-4.3-1.4-0.7,0.8-0.7,1.4c-0.1,3,4.3,5.1,2.5,7.9c-1.2,1.9-0.7,3.2-0.6,5.2    c-4-1.6-4.1-5.8-6.5-8.4c-1.3,6.4,7,8.3,6.4,14.8c-3.2-3.1-3.9-7.4-7.7-9.9c1.8,6.2,6.3,10.4,9.4,16.3c-4.8-1.6-4.1-6.4-7.6-8.5    c0.4,7.8,8,11.6,9.6,19.1c-2-1.7-3.6-2.8-4.4-4.6c-0.5-1.2-1.4-1-2-0.7c-1.2,0.7-0.2,1.3,0.3,2c2.1,2.6,4.1,5.3,6.2,7.8    c0.6,0.8,0.9,1.3,0.1,2.1c-1,1.1-1.4,0-1.9-0.5c-1.2-1.1-2.4-2.2-4.2-2.8c1,2.1,2.1,4.2,3.4,7c-7.1-2-11.4,1.2-14.7,6.3    c-0.8,1.3-2.1,1.8-3.4,1.7c-1.9,0-4.1,0.3-5,1.4c-3.2,3.6-4.9,1.5-7-1.4c4-1.2,7.2-3.5,10.2-6.3c4.9-4.6,7.3-9.5,8.2-16.6    c1.3-10.4,1.1-21.2,4.2-31.3c2.6,1.8,2.5,6,6.1,7.3c-0.4-4.1-4-6.5-4.3-10.7c-0.1-2,0.9-2,1.4-2.9c0.2-0.1,0.3-0.2,0.3-0.2    C102.1,179.6,102,179.7,101.8,180z"
		pathData3 := "M225.4,227.9c-1.5-1.2-3.6-1.9-2.3-4.3c4.8,3,11.1,4.3,14.3,2.7c-2.2-0.8-3.5-2.6-6.3-1.5    c-1.9,0.7-4.5-1-5.3-2.7c-4.7-9.3-10.7-17.7-16.2-26.5c-1.5-2.3-2.3-5.6-1.9-9.4c5.2,6.1,9.3,12.4,14.8,17.5    c2.4,2.2,4.8,4.4,7.4,6.4c0.6,0.5,1.8,1.9,2.7,0.4c0.7-1.2-0.7-1.2-1.4-1.8c-8.3-6-15.1-13.5-20.9-21.8c-0.4-0.6-1.6-1.4-0.6-2.8    c3,3.4,6,6.7,8.9,10c0.6,0.7,0.6,2.1,1.9,1.2c1.6-1.2-1.2-1.5-0.5-2.8c2.4,0.6,2.8,3.3,5.8,4.1c-2.4,1.1-2.7-0.9-4.1-0.5    c0.2,3.6,4.3,4,5.7,6.6c1.1,2.1,3.6,3.4,5.4,5.2c0.7,0.7,1.4,0.9,2.3,0.7c0.3-0.1,0.7-0.6,0.6-0.9c-0.2-0.8-0.8-1.4-1.7-1.6    c-3.5-0.8-4.6-4.1-6.6-6.4c4.6,2.4,8.9,5.3,10.9,10.4c0.7,1.7-0.3,2.5-2,1.8c-0.9-0.4-1.7-1-2.8-0.3c1.3,2.7,4.8,3,7.2,6.4    c-15.1-5.1-23.2-16.2-31.2-27.6c1.7,5.1,4.6,9.5,8.1,13.6c3.4,3.9,7.6,6.8,11.5,10.2c4.2,3.7,9.5,4.8,14.3,7.6    c-3.6,0.9-6.9-0.2-11.9-2.6c-3.8-1.9-7-4.4-9.6-7.8c3.6,9.4,8.6,10.7,18.4,13.7c0.8,0.3,2.1-0.6,2.3,1.1c0.2,1.7-1.3,1.9-2.1,2.1    c-3.2,0.7-6.4,1.8-9.8,0.8c-1.9-0.6-3.7-0.4-5.5,1c-3.1,2.4-5.7,5.4-9.3,7c-1.4,0.6-2,3.2-4.1,1c-0.6-0.6-2.4,0.7-2.5-1.3    c-0.1-1.7,1.2-2.1,2.4-2.5C216.3,232.6,220.7,230.4,225.4,227.9z"
		pathData4 := "M53.4,89.1c3.6-0.7,6.7-2.3,9.5-4.4c1-0.8,2-1.3,3.2-1.7c0.9-0.3,1.7-0.9,2.7-0.3c-0.6,0.8-0.3,2-1.8,2.4    C60.3,87,55.5,91.8,50.8,97c-1.4-9.3,1.4-15.6,9.4-18.2c1-0.3,1.9-1.2,2.9-1.5c1.6-0.4,3.5-3,4.7-1.3c1.6,2-1.1,3.7-2.3,5.1    c-1.2,1.5-3.2,2.9-5,3.4C57.6,85.3,55.5,87.3,53.4,89.1z"
		pathData5 := "M220.6,226.2c-1,3.3-2.9,4.1-4.9,4.7c-2.4,0.8-5.1,1.2-7.4,2.2c-1.9,0.9-2.5,0-3.1-1.4    c-1-2.4,1.6-2.2,2.2-2.4c2.7-0.4,5.5-0.2,8.3-0.9C217.2,228.2,218.7,227.8,220.6,226.2z"
		pathData6 := "M27.5,155.5c-0.2-4.1-0.8-7.5,2.1-10.1c0.8-0.7,1.8-1.1,2.6-0.3c0.8,0.8-0.3,1.5-0.8,2    C28.9,149,29.8,152.5,27.5,155.5z"
		pathData7 := "M82.9,234.5c-3.4,2.5-3.4,2.5-9.9,2.3C74.8,234.2,74.8,234.2,82.9,234.5z"
		pathData8 := "M198.9,241.4c1.6-1.5,3.3-3,4.9-4.6c0.8-0.8,1.8-1.1,2.5-0.3c0.7,0.8,0.3,1.8-0.5,2.6    C204,240.8,201.4,240.8,198.9,241.4z"
		pathData9 := "M219.1,243c1.7-1.5,3.4-2.9,5-4.4c0.8-0.8,1.8-1.2,2.5-0.6c1,0.9,0.4,2-0.5,2.8    C224.2,242.5,221.8,242.8,219.1,243z"
		pathData10 := "M69.8,247.2c0.7-2.3,1.6-3.6,2.9-4.9c0.9-0.9,1.7-1.1,2,0.1c0.1,0.6,2.3,1.4,0.1,2.4    C73.3,245.4,72.2,247,69.8,247.2z"
		pathData11 := "M45.6,139.4c-2.3,3.2-5.2,4.8-8.5,5.4c-0.2-0.4-0.3-0.7-0.5-1.1C39.5,142.3,42.5,140.8,45.6,139.4z"
		pathData12 := "M95.6,248.1c2.2-6.8,3.2-7.5,5.9-4.7C99.6,244.9,97.8,246.3,95.6,248.1z"
		pathData13 := "M60.8,244.3c1.5-1.5,2.9-3.1,4.4-4.6c0.5-0.6,1.1-1.4,2-0.8c0.4,0.3,0.9,0.8,1,1.2c0.2,1-0.7,1.8-1.4,1.8    C64.6,242.1,63,243.7,60.8,244.3z"
		pathData14 := "M104.3,247.9c-0.6-2,0.4-3.6,1.7-4.4c0.9-0.6,2.4,0.4,3.2,1.6C107.9,246.6,106,247,104.3,247.9z"
		pathData15 := "M239.8,242.6c-2.4,0.7-3.3,3.3-5.8,3.5C236.5,241.1,237.5,240.4,239.8,242.6z"
		pathData16 := "M87.1,246.5c1-2.3,1.6-4.4,4.1-4.4c0.3,0,0.8,0.3,0.9,0.6C92.4,243.6,90.9,244.9,87.1,246.5z"
		pathData17 := "M227.5,245.5c0.1-3.9,0.5-4.2,4.5-4.1C231,243.5,228.3,243.3,227.5,245.5z"
		pathData18 := "M70.2,241.9c-1.7-6.2,3.4-1.6,4.7-2.9C73.4,239.9,72,240.8,70.2,241.9z"
		pathData19 := "M212,239.7c-2.1,1.6-3.7,2.7-5.4,4C206.6,240.2,208.3,239.1,212,239.7z"
		pathData20 := "M248.4,223.2c0-0.9,0.2-2.1,1.3-1.5c1.7,0.9,1.6,3,2,4.7c-0.5-0.1-1-0.3-1.5-0.4C249.6,225,249,224.1,248.4,223.2z"
		pathData21 := "M220.1,239.4c1-2.5,1.6-4,3.2-4.7c0.2-0.1,0.9,0.6,0.9,0.8C223.6,237,222.3,238,220.1,239.4z"
		pathData22 := "M199.6,237.1c1.2-1.8,1.3-3.7,3.4-4c0.2,0,0.8,0.8,0.8,0.8C202.7,235,202,236.7,199.6,237.1z"
		pathData23 := "M66.7,236.7c1.5-1.1,2.8-2,4.7-1.1C70.2,236.7,68.8,237.3,66.7,236.7z"
		pathData24 := "M19.9,136.9c0.6,0.7,1.1,1.2,1.3,1.7c0.1,0.2-0.3,0.9-0.5,0.9c-0.8,0.3-1.3-0.4-1.7-1    C19,138.3,19.5,137.7,19.9,136.9z"
		pathData25 := "M66.1,106.9c-1.3,1.7-2.6,0.9-3.6,1.4c-1.6,0.8-3.8,0.7-2.1,3.8c1.3,2.5-0.2,6-2.8,6.6c-2.6,0.6-6.4,2.4-8.5-1.9    c-0.6-1.1-2-3.7-4-0.9c-0.7,0.9-1.1-0.3-1.5-0.6c-0.6-0.5-1.5-1.1-0.7-2.1c0.5-0.6,1.8-1.6,1.8-1.6c3,3,4.9-0.6,7.2-1.3    c3-0.9,5.7-2.6,8.6-3.8C62.3,106.1,64,106.1,66.1,106.9z"
		pathData26 := "M86.8,95.4c-2.3,1.2-5.3,0.4-4.6,4.1c0.2,0.9-2.4,1.9-3.8-0.3c-3-5.1-0.4-11.7,5.2-13.4c1-0.3,2.1-0.9,2.9-1.6    c1.5-1.2,2.4-1.6,3.5,0.5c0.9,1.8,0.6,3.2-1.3,3.5c-1.7,0.3-2.4,1.4-3.4,2.3C83.2,92.4,81.9,94.2,86.8,95.4z"
		pathData27 := "M160.2,159.2c-1.2,3.3-3.2,5-6.6,3.9c-2.1-0.7-4.1-0.7-6,0.4c-0.7,0.4-1.8,1-2.3,0.8c-3.1-1.5-5.8,0-8.6,0.9    c-0.5,0.1-1,0-2.1-0.1c0.8-0.8,1.2-1.4,1.7-1.8c2.3-1.5,2.7-3.1,0.5-5.1c-0.7-0.6-1.8-1.2-0.7-2.4c0.4-0.4,1-0.9,1.4-0.8    c1,0.2,1.5,1,1.6,2.2c0.2,3.1,3,3.6,5,4.8c0.2,0.1,0.4,0.3,0.6,0.4c0.5-4.8,5.2-0.2,6.6-2.9c0.1-0.2,1-0.2,1.1-0.1    C155.1,163.2,157.5,161.2,160.2,159.2z"
		pathData28 := "M202.7,141.3c0.4-2-2.9-3.8-0.4-4.9c2-0.9,1.8,2.4,2.4,3.7c1.7,3.7,2.5,4.2,6.3,2.3c0.6-0.3,1-0.2,1.5,0.1    c2.5,1.7,5.1,3.3,7.4,5.1c0.3,0.2,1.6,1.4,0.5,2.4c-1.1,1-1.3-0.3-1.8-0.7c-2.2-1.7-3.6-5.2-7.6-2.3c-1.9,1.4-3.7-1.8-5.8-2.4    c-2.8-0.8-5.2-2.4-7.2-4.6c-0.5-0.6-1.5-1.3-0.3-2c0.6-0.3,1.6-0.6,2,0.6C200.3,140,201.1,141.1,202.7,141.3z"
		pathData29 := "M132.6,218.2c-0.3-1.9-0.5-3.6-1.5-5.2c-0.3-0.5-0.6-1.8-0.4-1.9c3.6-2.2,0.5-3.9-0.4-5.6c-0.4-0.8-0.8-1.6-0.3-2.3    c0.7-1.1,1.4,0.1,2.1,0.2c2.6,0.4,4,2.2,4.5,4.7c0.4,2.1-2.6,1.9-2.5,3.7C134.4,214,134.1,215.8,132.6,218.2z"
		pathData30 := "M133.6,180c-5,1.5-12-1.5-13.9-5.7c-0.6-1.2-0.2-2.2,0.9-2.9c1.3-0.7,1.9,0.4,2.2,1.2c0.7,1.5,1.4,3,2.6,4.1    c1,0.9,1.9,2.3,2.9-0.2c0.5-1.3,1.7-1.8,2.5,0.2C131.3,177.9,132.5,178.7,133.6,180z"
		pathData31 := "M109.4,164.3c4,7.9,10.3,12.3,16.1,17.5C120.3,182,110,171.2,109.4,164.3z"
		pathData32 := "M261.1,166.9c0.1,3.8,0.4,6.8-1.3,10c-1.8,3.4,2,5.8,3.8,8.6c-4-0.4-7.4-4.8-5.6-8.6c1.3-2.7,1.1-5.3,1.7-7.9    C259.9,168.3,260.5,167.8,261.1,166.9z"
		pathData33 := "M137.7,100.7c-1.2-3.1,2.1-1.4,2.7-2.6c2,1.3,2.5-0.5,3.8-1.8c-2.2,0-4,0-5.8,0l0.1,0.1c3.6-2,1.5-3.5,10,0.9    c-3.5-0.1-5.5,4.8-9.7,2.4c-0.2-0.1-0.8,0.6-1.3,0.9L137.7,100.7z"
		pathData34 := "M263.8,230c-4.8,1.5-9.7-2.1-11-8.3C256.7,224.9,259.2,228.9,263.8,230z"
		pathData35 := "M55,131.5c-1.3,0.7-1.9-0.3-2-1.2c-0.1-1.6,2.3-2.3,2.5-2.2c2.1,1.7,3.6,0,5.1-0.9c1-0.6,1.6-0.3,2.1,0.5    c0.6,1.1-0.4,1.5-1.1,1.6C59.3,129.6,57.2,130.6,55,131.5z"
		pathData36 := "M132.6,189.1c-4.2,0.2-6.6-3.8-10.7-3.4C124.7,183.7,129.4,185.2,132.6,189.1z"
		pathData37 := "M188.1,101.5c1.5-1.9,3.9-2.4,5.8-1.8c1.5,0.5,3.2,0.2,4.6,0.9c-3.3,2.6-6.8,1.7-10.4,0.8L188.1,101.5z"
		pathData38 := "M95.4,112.1c-1.5,1.1-2.8,1.8-3.6,2.8c-1.8,2.4-3.1,1.5-4.5-0.2c-0.8-0.9-1.7-2.2-0.7-2.8c1.2-0.8,2.6-0.3,1.6,1.7    c1,0.5,2.2,2.1,3.1,0C91.9,111.8,92.7,111,95.4,112.1z"
		pathData39 := "M133,196.4c-2.3-1.9-4.8-2.6-6-5.2c-0.2-0.5-0.8-0.9-0.1-1.4c0.3-0.2,0.9-0.5,0.9-0.4C129.5,191.6,133.2,192.1,133,196.4z    "
		pathData40 := "M96.8,118.6c-4.2-2-5.2,5.7-10.1,1.2C90.7,121.3,91.7,113.5,96.8,118.6z"
		pathData41 := "M230.7,140.6c-0.2-1.1-0.3-1.9-0.6-2.7c-0.3-1.1-1.3-2.5,0.2-2.9c1.5-0.4,2,1.5,2.6,2.7c0.6,1.2-0.8,2-0.8,3.1    c0.1,1.2-0.6,1.5-1.6,1.3C228.8,141.8,231,140.9,230.7,140.6z"
		pathData42 := "M240.2,190.8c-4-1.5-6-3.6-5-8.6C237.1,185.4,238.5,187.8,240.2,190.8z"
		pathData43 := "M266.3,227.4c-3.2,2.2-5,0.2-6.7-2.1c-0.4-0.5-0.4-1.5,0.3-1.9c0.8-0.4,1,0.2,1.4,0.9    C262.3,226.4,264.3,226.8,266.3,227.4z"
		pathData44 := "M137.9,93.1c-4,0.6-6.9,3.6-10.3,5.4C128.4,95.7,133.3,93,137.9,93.1z"
		pathData45 := "M132,91.5c-3.8-0.3-6.4,2.2-9.5,3.4C123.6,92.2,129.1,89.9,132,91.5z"
		pathData46 := "M241.8,183.2c1.7,0.6,1.6,3.4,3.8,3.8c-0.4,0.7-1,1.5-1.9,1.3c-1.9-0.5-2.5-2.2-3-3.9    C240.6,183.9,240.8,183.2,241.8,183.2z"
		pathData47 := "M208.7,154.1c3.3,0.2,3.8,3.3,6.1,3.6c0.1,0,0.4,0.9,0.3,1.1c-0.4,0.7-1.4,0.7-1.8,0.4C211.5,158,210,156.6,208.7,154.1z"
		pathData48 := "M258.1,185.9c2.5,1.5,5.2,2.8,5.4,6.3C262.4,189.6,258.6,189.2,258.1,185.9z"
		pathData49 := "M132.7,202.9c-1.5-2-2.6-3.4-3.7-4.6c-0.6-0.7-1.1-1.2-0.1-1.9c0.7-0.5,1.5-0.6,1.7,0.2    C131.2,198.3,133.7,199.2,132.7,202.9z"
		pathData50 := "M55.1,122.8c-4.2,1.7-5.9-1-7.7-3.3c2.1-0.8,2.4,0.9,3.3,2C51.6,122.4,53.5,121.1,55.1,122.8z"
		pathData51 := "M118,92c-2.4,1.4-4.9,2.8-8.3,4.8C112,93.1,115.1,92.7,118,92z"
		pathData52 := "M259.7,194.6c0.2-0.1,0.3-0.2,0.4-0.2c2.4,0,4.6,1.1,5.9,2.9c0.8,1.1-0.4,1.6-1.5,0.8c-1-0.9-1.9-2-3.3-2.4    C260.8,195.4,260.3,195,259.7,194.6z"
		pathData53 := "M264.3,138.1c-0.2,2.7-1,5.1-4.2,6.7C260.9,141.8,263.1,140.3,264.3,138.1z"
		pathData54 := "M119.8,93.8c1.5-2.9,4-4.3,7.4-3.9c-2.5,1.3-5,2.6-7.5,3.8L119.8,93.8z"
		pathData55 := "M212.4,162c-2.7,0.4-4.9-0.3-5.5-3.3C208.7,159.8,210.5,160.9,212.4,162z"
		pathData56 := "M235.7,106.3c2-2.2,4.2-1.7,6.7-1.5C240.5,107.5,238.1,105.9,235.7,106.3z"
		pathData57 := "M123.4,106.2c1.7-1.7,1.2-4.8,4.6-6.2C126.9,102.9,125.8,105,123.4,106.2z"
		pathData58 := "M264.5,224.3c2.1,0,3.9,0,6.5,0C268.5,226.4,266.7,226.5,264.5,224.3z"
		pathData59 := "M216.3,149.8c0.7,1.3,2.8,1.5,1.9,3.4c-0.6,1.3-3.1,0.7-2.6,0.3C217.1,152.4,215.8,151.1,216.3,149.8z"
		pathData60 := "M106.7,110.4c0.5-0.6,0.8-1.2,1.1-1.2c1.6,0,2-2.4,3.8-1.8c0.2,0.1,0.3,0.7,0.2,0.7C110.2,108.9,109.3,111.1,106.7,110.4z    "
		pathData61 := "M198.9,102.6c1.5-1.7,3.2-2.5,5.5-2.2c-0.4,1.4-1.9,0.4-2.4,1.5C201.4,103,200.3,103.2,198.9,102.6z"
		pathData62 := "M132.2,224.1c-1-0.6-2-1.3-2.3-2.5c-0.2-0.7,0.6-1.1,1-0.8c0.9,0.6,1.6,1.4,2.3,2.3C133.4,223.5,133,224.1,132.2,224.1z"
		pathData63 := "M148,100.4c2.3-2.4,2.3-2.4,5.8-2.4C151.8,98.8,149.9,99.6,148,100.4z"
		pathData64 := "M157.1,98.4c1.1-1.6,2.4-1.9,3.7-2.2c0.1,0,0.5,0.8,0.4,1C160.5,99,158.9,98.8,157.1,98.4z"
		pathData65 := "M135.3,183c-0.1-1.3-1.9-2.4-0.3-3.8c0.3-0.2,1-0.4,1.1-0.2C137.5,180.6,135.5,181.6,135.3,183z"
		pathData66 := "M254.5,178.8c-1.6,2.5-3,1.7-5,1.6C251.2,178.3,252.7,179.2,254.5,178.8z"
		pathData67 := "M137.6,100.6c-1.6,0.8-3.3,1.7-4.9,2.5c-0.2-0.4-0.3-0.8-0.5-1.2c1.8-0.5,3.1-2.1,5.5-1.2    C137.7,100.7,137.6,100.6,137.6,100.6z"
		pathData68 := "M83.4,137.4c-1.3,1.1-2.6,1.8-4.5,1.2C80.2,137.1,81.6,136.7,83.4,137.4z"
		pathData69 := "M138.5,96.3c-1.8,0.2-2.5,3.1-5.3,2.1c1.9-0.9,3-3.1,5.4-2C138.6,96.4,138.5,96.3,138.5,96.3z"
		pathData70 := "M248.4,108.1c-0.5,1.1-1.4,0.7-2.2,0.7c-0.8,0.1-1.5-0.4-1.1-1.2c0.4-0.8,1.5-0.5,2.3-0.5    C248.1,107.2,248.3,107.6,248.4,108.1z"
		pathData71 := "M82.7,114.3c-0.8,1-1.3,2.3-2.8,2.2c-0.2,0-0.5-0.8-0.4-1.1C80.2,114.2,81.2,113.6,82.7,114.3z"
		pathData72 := "M260.9,160.4c-0.1,2-1.6,2.8-2.9,4.1C257.6,161.9,260.2,161.7,260.9,160.4z"
		pathData73 := "M101.6,121.8c-0.2-1.5,0.4-2.4,1.9-2.6c0.2,0,0.5,0.4,0.5,0.5c-0.1,1.5-1.1,2.1-2.5,1.9    C101.5,121.7,101.6,121.8,101.6,121.8z"
		pathData74 := "M120.8,114.1c-0.6,1.7-2,2-3.5,2.7C117.6,114.3,119.8,115.2,120.8,114.1z"
		pathData75 := "M188,101.4c-0.1,0.1-0.2,0.3-0.3,0.3c-1.4,0-2.5,1.6-4,0.7c-0.1-0.1-0.1-0.7,0-0.8C185.1,100.9,186.6,101.1,188,101.4    C188.1,101.5,188,101.4,188,101.4z"
		pathData76 := "M263.7,147.2c0.1-1.2,0.6-2.2,1.9-2.3c0.2,0,0.8,0.5,0.7,0.7c-0.4,1.4-1.2,2.2-2.7,1.5L263.7,147.2z"
		pathData77 := "M248.3,181.9c-0.6,0.3-1.1,0.7-1.7,0.9c-1.2,0.3-0.9-1-1.2-1.5c-0.5-0.9,0.6-0.5,0.9-0.6    C247.3,180.4,247.7,181.3,248.3,181.9z"
		pathData78 := "M84.4,129.7c-0.8,1.4-1.7,1.7-3,1.2C81.9,129.4,82.9,129.1,84.4,129.7z"
		pathData79 := "M119.7,93.7c-1,0.8-1.9,1.6-3.2,2.6c0.2-2.5,1-3.4,3.3-2.5C119.8,93.8,119.7,93.7,119.7,93.7z"
		pathData80 := "M121.9,111.7c-1,0.6-1.7,1.5-2.8,1.3c-0.2,0-0.5-0.9-0.3-1.1C119.6,111,120.7,111.5,121.9,111.7z"
		pathData81 := "M263.5,147.1c0.4,1.5,0.9,3.2-2,2.7c0.2-1.3,1.3-1.9,2.1-2.6C263.7,147.2,263.5,147.1,263.5,147.1z"
		pathData82 := "M168.6,97.3c-1,0.2-1.5,1.9-2.8,1c-0.2-0.1-0.3-0.6-0.2-0.7C166.5,96.4,167.6,97,168.6,97.3z"
		pathData83 := "M210.8,101.5c-1.1,1.7-2.4,1.7-4.7,0.8C208.2,102.4,209.2,101,210.8,101.5z"
		pathData84 := "M146.1,92.4c-0.8,0.1-1.2,0.2-1.7,0.2c-0.7,0-1-0.4-0.9-1c0.1-0.6,0.7-0.8,1.2-0.5C145.1,91.4,145.5,91.8,146.1,92.4z"
		pathData85 := "M264,151.9c-0.4,2.3-0.9,3.1-2.8,3C262,153.9,262.7,153.2,264,151.9z"
		pathData86 := "M228,102.1c1.3,0,1.7,0.3,1.6,1c-0.1,0.7-0.7,0.8-1.2,0.7c-0.3-0.1-1.2,0.3-0.9-0.6C227.6,102.6,228,102.2,228,102.1z"
		pathData87 := "M101.8,180c0-0.9-2-1.5-0.6-2.7c0.3-0.3,1-0.1,1.2,0.6c0.2,0.8,1.5,2-0.7,1.9C101.7,179.8,101.8,180,101.8,180z"
		pathData88 := "M55.6,117.2c-2.5,0.5-4.5,0.6-4.3-2.4c0.2-2.9,2.5-3.6,4.9-4.1c0.3-0.1,0.9,0.3,1,0.5    c0.2,0.9-0.6,0.9-1.2,1.2C53.2,113.5,51,114.9,55.6,117.2z"
		canvas.Path(pathData0, "fill:"+lineColor)
		canvas.Path(pathData1, "fill:"+hexCode)
		canvas.Path(pathData2, "fill:"+hexCode)
		canvas.Path(pathData3, "fill:"+hexCode)
		canvas.Path(pathData4, "fill:"+hexCode)
		canvas.Path(pathData5, "fill:"+hexCode)
		canvas.Path(pathData6, "fill:"+hexCode)
		canvas.Path(pathData7, "fill:"+hexCode)
		canvas.Path(pathData8, "fill:"+hexCode)
		canvas.Path(pathData9, "fill:"+hexCode)
		canvas.Path(pathData10, "fill:"+hexCode)
		canvas.Path(pathData11, "fill:"+hexCode)
		canvas.Path(pathData12, "fill:"+hexCode)
		canvas.Path(pathData13, "fill:"+hexCode)
		canvas.Path(pathData14, "fill:"+hexCode)
		canvas.Path(pathData15, "fill:"+hexCode)
		canvas.Path(pathData16, "fill:"+hexCode)
		canvas.Path(pathData17, "fill:"+hexCode)
		canvas.Path(pathData18, "fill:"+hexCode)
		canvas.Path(pathData19, "fill:"+hexCode)
		canvas.Path(pathData20, "fill:"+lineColor)
		canvas.Path(pathData21, "fill:"+hexCode)
		canvas.Path(pathData22, "fill:"+hexCode)
		canvas.Path(pathData23, "fill:"+hexCode)
		canvas.Path(pathData24, "fill:"+hexCode)
		canvas.Path(pathData25, "fill:"+lineColor)
		canvas.Path(pathData26, "fill:"+lineColor)
		canvas.Path(pathData27, "fill:"+lineColor)
		canvas.Path(pathData28, "fill:"+lineColor)
		canvas.Path(pathData29, "fill:"+lineColor)
		canvas.Path(pathData30, "fill:"+lineColor)
		canvas.Path(pathData31, "fill:"+lineColor)
		canvas.Path(pathData32, "fill:"+lineColor)
		canvas.Path(pathData33, "fill:"+lineColor)
		canvas.Path(pathData34, "fill:"+lineColor)
		canvas.Path(pathData35, "fill:"+lineColor)
		canvas.Path(pathData36, "fill:"+lineColor)
		canvas.Path(pathData37, "fill:"+lineColor)
		canvas.Path(pathData38, "fill:"+lineColor)
		canvas.Path(pathData39, "fill:"+lineColor)
		canvas.Path(pathData40, "fill:"+lineColor)
		canvas.Path(pathData41, "fill:"+lineColor)
		canvas.Path(pathData42, "fill:"+lineColor)
		canvas.Path(pathData43, "fill:"+lineColor)
		canvas.Path(pathData44, "fill:"+lineColor)
		canvas.Path(pathData45, "fill:"+lineColor)
		canvas.Path(pathData46, "fill:"+lineColor)
		canvas.Path(pathData47, "fill:"+lineColor)
		canvas.Path(pathData48, "fill:"+lineColor)
		canvas.Path(pathData49, "fill:"+lineColor)
		canvas.Path(pathData50, "fill:"+lineColor)
		canvas.Path(pathData51, "fill:"+lineColor)
		canvas.Path(pathData52, "fill:"+lineColor)
		canvas.Path(pathData53, "fill:"+lineColor)
		canvas.Path(pathData54, "fill:"+lineColor)
		canvas.Path(pathData55, "fill:"+lineColor)
		canvas.Path(pathData56, "fill:"+lineColor)
		canvas.Path(pathData57, "fill:"+lineColor)
		canvas.Path(pathData58, "fill:"+lineColor)
		canvas.Path(pathData59, "fill:"+lineColor)
		canvas.Path(pathData60, "fill:"+lineColor)
		canvas.Path(pathData61, "fill:"+lineColor)
		canvas.Path(pathData62, "fill:"+lineColor)
		canvas.Path(pathData63, "fill:"+lineColor)
		canvas.Path(pathData64, "fill:"+lineColor)
		canvas.Path(pathData65, "fill:"+lineColor)
		canvas.Path(pathData66, "fill:"+lineColor)
		canvas.Path(pathData67, "fill:"+lineColor)
		canvas.Path(pathData68, "fill:"+lineColor)
		canvas.Path(pathData69, "fill:"+lineColor)
		canvas.Path(pathData70, "fill:"+lineColor)
		canvas.Path(pathData71, "fill:"+lineColor)
		canvas.Path(pathData72, "fill:"+lineColor)
		canvas.Path(pathData73, "fill:"+lineColor)
		canvas.Path(pathData74, "fill:"+lineColor)
		canvas.Path(pathData75, "fill:"+lineColor)
		canvas.Path(pathData76, "fill:"+lineColor)
		canvas.Path(pathData77, "fill:"+lineColor)
		canvas.Path(pathData78, "fill:"+lineColor)
		canvas.Path(pathData79, "fill:"+lineColor)
		canvas.Path(pathData80, "fill:"+lineColor)
		canvas.Path(pathData81, "fill:"+lineColor)
		canvas.Path(pathData82, "fill:"+lineColor)
		canvas.Path(pathData83, "fill:"+lineColor)
		canvas.Path(pathData84, "fill:"+lineColor)
		canvas.Path(pathData85, "fill:"+lineColor)
		canvas.Path(pathData86, "fill:"+lineColor)
		canvas.Path(pathData87, "fill:"+lineColor)
		canvas.Path(pathData88, "fill:"+hexCode)
	},
}
