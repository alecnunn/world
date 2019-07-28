package charge

import svg "github.com/ajstarks/svgo"

var TigerPassant = Charge{
	Identifier: "tiger-passant",
	Name:       "tiger passant",
	Noun:       "tiger",
	NounPlural: "tigers",
	Descriptor: "passant",
	SingleOnly: false,
	Tags: []string{
		"aggressive",
		"animal",
		"cat",
		"strong",
		"tiger",
	},
	Render: func(canvas *svg.SVG, hexCode string, lineColor string, canvasWidth int, canvasHeight int, number int) {
		pathData0 := "M136.8,135.5c0.8,3.2,0,6.5-2.8,7.8c-6.1,2.8-1.6,8.4-3.6,12.2c-0.6,1.2,2.7,2.2,4.5,3    c1.2,0.5,1.9,0.5,2.5-0.4c1-1.4,1.1-2.5-0.9-3.4c-0.9-0.4-0.6-1.5,0.5-1.6c1.6-0.3,2.5,0.5,3,2c1.2,3.8-4.2,9-8.3,7.8    c-2.1-0.6-4.2-1.7-6.5-2.7c0,4.9,3.1,7.4,7,8.7c3.1,1,5.5-0.1,5.4-5.2c1.8,4.8,1.8,4.8-8.5,9.2c14.3,0.4,28-0.2,41.7-1.8    c18-2.1,35.8-2,53.2,4.3c7,2.5,14.1,0.7,16.3-8.1c1.8-7-3.7-11.1-8.4-10.2c-2.2,0.4-4.7,1.1-3.2,4.4c0.3,0.8,0.6,1.7-0.2,2.3    c-0.9,0.6-1.8,0.3-2.7-0.3c-1.9-1.1-3.8-1.8-3.9-4.5c-0.2-3.5,0.4-6.6,3.5-8.7c1.6-1.1,3.4-1.7,5.5-2.7c-2.4-2.9-5.9-1.7-8.6-2.3    c-6.8-1.5-14-1.3-20.9-2.3c-10.4-1.5-19.1-5.9-24-15.9c-4.4-8.9,0.7-22.7,13.5-24.8c4.1-0.7,8.4-1.4,12.7-1.3    c4.8,0.2,5.4-1.2,2.6-5.7c4.9-1.2,8,0.6,8.6,5.6c0.4,3,2.7,2.6,4.6,3.2c5,1.6,10.5-0.1,12.1-4.4c1.9-4.8-0.9-7.7-6.4-9.4    c5.8-1,9,1,10.8,5.1c1.6,3.7,1.9,7.3-1.6,10.7c-5.9,5.7-11.4,12.4-19.9,13.7c-5.5,0.8-9.2,6.2-15.1,5.2c-1.2-0.2-2.3-0.1-3.1-1.3    c0.8-1.7,4.5-1.9,3.2-4.4c-1.2-2.4-3.7-4.8-6.7-4.2c-2.7,0.6-5.6,2.3-5.9,5.8c-0.3,4.3,1.1,7.7,5.2,10.1c5.9,3.5,12.4,4.9,19,4.8    c4.8-0.1,9.6,0.6,14.2,1.6c3.4,0.7,3.5,0.1,3.6-2.7c0.1-5.3,2.8-8.5,7.9-9.7c5.2-1.3,10.7,0.7,13,4.7c0.9,1.6,1.6,3.2,0.7,5    c-1.1,0.3-1.3-0.2-1.8-1c-1.2-2.3-2.8-5.2-5.8-3.2c-2.5,1.7-4.3,4.4-2.3,8c0.8,1.5,1.9,2.7,3.1,3.8c10.6,9.6,10.3,24.5,3.5,35.7    c-3.4,5.6-8.2,9.8-14.4,11.8c-1.7,0.6-2.4,1.1-1.9,2.5c1.2,3.4,1.1,7,2.1,10.4c0.5,1.8,0.2,4.6,2.9,5.1c2.6,0.5,4.8-0.1,6.1-3.1    c1.6,4.5-1,7-4,9.1c-1.1,0.8-2.3,0.8-1.3,2.7c3.2,5.8,4.6,12.3,10,17c3,2.7,4.8,2.2,7.3-0.1c1-1,1.6-3.4,3.8-2c1.8,1.1,1,3,1,4.6    c0,2-0.9,4.1-2.7,5c-2.8,1.4-4.1,3.3-3.2,6.3c1.1,3.5,3.9,0.4,5.8,0.9c0,0.6,0,1.1-0.1,1.2c-6.9,2.4-6.6,8.9-8.3,14.4    c-0.7,2.4-0.9,4.2,0.9,6.3c2.7,3,2,6.9-1.1,8.8c-2.8,1.7-3,1.3-4-1.9c-0.5-1.6-2.6-3.3-4.3-3.8c-1.6-0.5-2.1,1.5-2.1,3.1    c0,1.9-1.1,2.8-2.8,2.6c-1.3-0.1-2.1,0.6-3,1.3c-1.2,0.9-1.8,3.4-3.8,2.3c-2-1-1.2-3.2-1.2-5c0-0.7,0.3-1.5,0.9-3.5    c-4.4,5-9.5-0.1-13.8,2.8c-0.5,0.3-1.9,0.4-2-0.7c-0.2-1.1-0.4-2.2,0.4-3.3c3.2-4.3,8.3-6.2,12.1-9.8c0.5-0.4,1.9-0.5,2.4-0.1    c4.7,4.3,5.2-0.6,6.3-3c3.2-7,2.2-12.1-3.5-17.6c-2.4-2.3-4.3-5.2-7.8-6.6c-3.2-1.3-5.4-2.5-7.6,1.1c-1,1.7-3,1.5-4.4,0.5    c-2-1.4,0.4-2.1,0.9-3c1.8-3,1.3-5.6-1-8c-1.4-1.4-2.8-2.8-4.2-4.2c-3.9,6.8-3.8,9.1,0.2,14.4c2,2.5,4.1,4.6,7.5,2.1    c0.9,4.2-3,3.7-4.3,3.9c-2.6,0.4-3.9,1.3-4.2,3.7c-0.2,2.1-0.4,4.2,2.8,4.8c2.4,0.5,0.7,1.9-0.3,2.6c-2.5,1.8-5.4,1.1-8.1,1.3    c-0.8,0.1-2.4-2.2-2.6,0.8c-0.1,1.2-1.8,2.1-0.8,3.7c0.7,1,3,1.6,1.2,3.5c-1.5,1.6-3.5,0.6-5.2,0.9c-1.4,0.3-2.5-3.3-3.8-0.7    c-0.9,1.9-1,4.3,0.8,6.5c3.8,4.4,1.7,8.3-2.7,10c-1.4,0.5-3,0.6-2.5-1.5c0.8-3.6-2.5-4.3-4.1-4.9c-1.8-0.6-1.8,2-2,3.6    c-0.1,1.3-1,2-2.4,2.1c-1,0-1.8,0-2.8,0.7c-1.5,1-2.3,4.3-4.5,2.8c-2.6-1.7-1.2-4.6-0.5-8.2c-4.3,5-9.2,0.4-13.7,2.5    c-2.9,1.3-2.4-2.4-1.6-3.5c3.1-4.6,8.6-6.4,12.5-10.2c0.3-0.3,1.5-0.4,1.9,0c5.6,4.4,6.4-1.5,8.1-4.1c2.2-3.3,4-6.9,6.6-10.1    c3.3-4,3.2-11,0.2-15.5c-0.6-1-1.3-1.7-2.8-0.9c-2.1,1.2-4.6,0.2-6.8,0.5c-3.1,0.4-4.8-1.6-5.8-3.7c-0.9-2.1-0.6-4.7,2.3-5.9    c0.6,1.6,2,2.9,3.4,4.1c1.5,1.3,2.9,0.5,3.8-0.8c0.7-0.9,2.6-1.2,1.1-3.4c-2.6-4-3.4-8.7-2.9-13.7c-5,0.4-9.2,3.1-13.4,4.9    c-6.8,2.9-13.8,3.4-20.7,4.5c-6.5,1-12.4,4.9-19.3,4.3c-0.2,0-0.6,0.3-0.8,0.5c-3,4.7-5.4,9.6-6.8,15c-0.2,0.7,0.1,1.9,0.6,2.3    c5.9,4.3-0.7,4.5-2,4.8c-2.1,0.4-2,1.4-2.8,2.9c-2.7,5.4-3.8,11.5-8.1,16.1c-2.5,2.7-6.5,5-1.4,9.4c2.8,2.4-0.7,8.5-4.6,8.7    c-0.8,0-1.5,0.1-1.8-1.1c-1.3-4.3-1.4-4.3-5.7-7.1c0.8,6.9-7.1,4.7-9.2,9c-0.2,0.4-2-0.9-2-2.1c0-1.9,0.2-3.8,0.3-5.9    c-2.9,1.2-5.7,3.9-8.8,0.5c-0.2-0.2-1.8,1-2.8,1.4c-1.2,0.4-2.9,1-3.5-0.4c-0.5-1-0.3-2.4,0.7-3.7c3.2-4.1,8.3-5.7,11.9-9.4    c0.5-0.5,2.3-0.4,2.8,0.1c3.6,3.6,4.7,0.3,5.8-2.2c5-10.5,8.2-21.6,10.6-32.9c0.3-1.2,0-2.6,0-3.8c-8.1,1-15.7,0.8-22.3-4.3    c-1.5,1-0.3,2.8-1.2,4.1c-1.1,1.4-1.6,3.7-3.5,3.7c-1.3,0-1.9-2-2-3.4c-0.1-2-1.6-1.4-2.6-1.3c-2.7,0.2-0.1,1.5-0.4,2.4    c-0.7,2.3,0.3,4.3,2.5,5.4c0.8,0.4,1.7,0.6,1.3,1.6c-0.3,0.7-1.1,1.5-1.8,1.5c-3,0.1-6.5,1.1-8.1-2.9c-0.7-1.6-2.1-2.5-3.3-3.7    c-1.7,0.7,0.5,3.1-1.9,3.8c-4.2,1.2-3.2,3,0.4,4.7c-3.3,1.6-6.1,2.3-8.4-0.5c-1.9-2.3-1.2-9.2,1-11.8c-0.3-0.1-0.7-0.4-0.9-0.3    c-1.8,0.6-3.4,3.7-5.3,0.9c-1.5-2.3-0.7-4.6,1.4-6.5c2.1-1.9,3.8-4.1,7.2-3.4c2.4,0.5,3.5-1,4.3-3c3.5-7.9,9.5-11.6,18.2-11.1    c5.3,0.3,10.6,0.1,15.9,0.1c0.9-2.4-2.1-2.9-1.9-4.9c-1.1,1.1-1.4,2.4-1.2,4.1c-2.7-1.4-2.5-4.2-3.4-6.1c-1.2-2.6-1.9-5.1-5.6-4.2    c-1.5,0.4-3.6-0.1-3.3-1.9c0.4-2.8,1.9,0.1,2.9,0c1.1-0.1,2.1-0.3,2.5-1.3c0.4-0.9-0.8-1.2-1.2-1.8c-1.5-2.5-4-2-6.2-1.9    c-4.3,0.1-8.4-0.4-12-3.2c-3.4-2.7-3.1-6.2-0.7-8.6c2.2-2.2,5.4-5.1,9.3-1.9c-0.3,0.3-0.4,0.4-0.5,0.4c-2.3,0.3-4.7-0.1-6.1,2.4    c-0.7,1.3,0.2,1.8,0.9,2.5c2.2,2.1,4.8,3.1,7.6,1.9c2.8-1.2,6-2.3,1.6-5.9c-2.1-1.7-3.9-4.3-1.4-7.2c2.6-3,5.7-4.9,9.8-2.5    c-0.7,1.6-2.2,0.5-3,1.5c-0.7,0.9-1.2,1.8-0.7,2.7c0.5,1.2,1.4,2.4,2.8,2.6c1.1,0.2,1.7-0.8,2.1-1.8c1.4-3.6,3.9-6.3,6.4-9.1    c3.3-3.7,2-6.6-2.8-8.1c-5.9-1.9-11.8-0.4-17.7-0.9c-1.9-0.2-3.8-0.3-5.7-0.8c-4-1.1-4.8-3.3-3.5-6.9c-6-0.6-7.7-2.3-7.1-7.6    c5.1,4.6,10.6,1.3,16,1.2c-0.4-1.1-0.3-3.3-1.8-2.7c-4.5,1.8-3.6-1.4-3.8-3.6c-0.2-4.6-3.1-6.2-7.2-3.7c-1.7,1-3.4,2.2-4.6,0.3    c-1-1.6,0.1-3.5,1.8-4.6c5-3.2,10-6.6,16.4-6.5c4.3,0,8.6,0.4,12.2-2.5c-1-3.3-7.4-4.2-5.1-9.3c0.9-2,2.9-4.3,6.2-2.8    c-4.7,4.5,0.7,5.7,2.8,8.3c0-3.6,0.5-7.1,2.5-10c2.3-3.4,6-2.7,9.5-2.5c-0.1,1.6-1.1,1.9-2.3,1.9c-3.2-0.2-3.4,2-3.2,4.3    c0.2,2.2,1.3,4.7,3.2,5.1c2.1,0.4,2.3-2.2,3.1-4c2.8-6.5,9.1-8.3,15-10.4c0.8-0.3,2,0.1,1.6,1.2c-1.8,4.9-1.2,10.2-2.5,15.5    c4.2-5,9.7-7.6,15.7-9.4c3-0.9,4-0.3,4.1,3c0,4-1.3,7.4-4.1,10c-1.6,1.5-2,3.1-0.4,4.3c1.4,1,3.5,2.5,5.1,0    c0.8-1.2,1.1-3.3,2.9-2.5c2.1,0.8,1.4,3.1,1.4,4.9c0,1.9-1.2,3.3-2.3,4.7c-1.8,2.4-4.7,4.6-0.9,7.8c1.8,1.6-1.7,2.9-1,4.7    C129.5,139.6,132.1,139.6,136.8,135.5z"
		pathData1 := "M104.1,212.1c-3.9-0.5-4.8-4-6.6-6.9c-1.7,2.6,3.3,5.1-0.3,7.6c-3-8.6-4.8-16.7,4.4-23.7    c-7.3,2.4-8.1,3.7-8.9,13.3c-4.8-10.1,3.9-12.9,8.6-17.6c-5.3,0.8-8.3,4.4-11,8.2c-1.7,0-1.1-1.6-1.9-2.2    c-0.9-0.7-1.6-1.1-2.5-0.1c-1.5,1.6,0.7,1.1,0.9,1.7c1,2.8,4.6,4.8,2.5,9.1c-3.6-6.1-8.7-11.5-0.2-18.2c-4.3,0.9-4.1,1-5.5,4.6    c-0.9,2.2,0.4,5.1-2.2,6.9c-2.3-6.8,2.3-11.3,5.5-16.5c-4.2,1.5-2.3-2-3-3.4c0.3,0.1,0.5,0,0.6-0.3c0,0.3-0.2,0.4-0.5,0.5    c-2.7,3-1.8,7.3-3.8,10.7c-0.7,1.2-1.1,2.1-2.9,2.5c1.1-3.5,2-6.6,3.1-10.2c-3.2,1.4-2.1,4.4-3.9,5.8c-6.3-4,3.1-5.9,1.2-9.3    c-10.5,8.4-17.4,8.8-21.4,0.8c2.4,2.2,5,3.6,8.3,3.8c3.6,0.2,5.8-1.8,8-4.2c1.4-1.6,4.1-1.3,5.4-4c-4.2,2.1-7.1,0.9-9-2.8    c-1.1-2.1-1.2-2.1,1.7-5.1c0.2,2.3,0.5,4.2,2.2,5.6c2.6,2.1,4.4,2.4,6.2-1c1.6-3.2,3.5-6.2,6.4-8.6c3.5-2.9,2.2-6.9,2.6-10.7    c1.8,1.1,1.8,3,2.2,4.6c5.7-3.2,6.8-5.7,4-10c-0.5,2.9-0.1,5.8-3,8c-0.2-1.6-0.2-3-0.5-4.3c-0.1-0.7-0.9-1.8-1.6-1.3    c-5.3,3.8-10.7-0.3-16.2,0.4c-4.8,0.6-9.8,1-14.7-0.7c-1.2-0.4-1.9-1-2.7-2.6c2.6-0.7,4.9-0.5,7.3-0.1c6.4,1.2,12.8,1.1,19-0.6    c3.1-0.8,4.8-3.6,4-7c-0.9-4.1-2.7-8.2-7.3-8.3c-1.6-0.1-3.2-1.1-5-0.3c0.9,2.2,3.2,1.6,4.7,3c-5.1-0.8-10.4-1-15.1-2.8    c-2.9-1.1-6.3-4.3-3.3-9.2c-2.8,3.7-6.7,4.2-11.4,4.7c5.3-4.4,10.7-6.7,16.7-6.7c5.4,0,10.1-1.2,15-3.6c6.2-3,12.8-6,19.6-1.2    c1,0.7,3.3,0.6,2.6,2.6c-0.5,1.5-2.3,0.9-3.5,1c-0.8,0.1-1.6,0.1-2.5,0.1c0.6,1.8,2.3,0.1,3.1,1.4c-0.4,1.1-3.1,1-2.1,3.5    c5.1-1.3,8.6-4.8,11.6-8.9c3.5-4.7,7.5-8.4,13.7-9.3c2.5-0.4,2.6,0.3,2.3,2.1c-0.1,0.6-0.4,1.4-0.9,1.6c-3.7,1.7-1,8.2-6.4,8.5    c1.1,5.3-4,5.7-6.5,8.1c-2.1,2-5,2-7.7,0.7c-1.1,2,1.2,2.4,1.3,3.8c-1.1,0.1-2.2,0.3-2.9,0.3c1.4,3.7,2.8,7.5,4.2,11.2    c-2.1,0.5-1.6-1.9-2.7-2.4c-0.4,0.3-0.9,0.5-0.9,0.6c1,3.2,1.2,6.8,4.4,9.2c-4.6,1.3-4.7-6.9-9.8-4c1.3,1.6,1,3.8,1,6.5    c1.7-1.9,3-3.4,5.1-1.3c2.1,2,4.2,0.8,6.1-0.7c1-0.8,1.7-1.9-0.5-2.3c-3-0.6-3.4-2.1-1.8-4.6c1.8-2.7,1-6.6,0.1-8.3    c-1.7-3.3,0.3-3.2,1.8-4.6c1.8-1.6,2.2-0.7,2.8,0.7c3.6,7.7,4.1,15.7,2.8,23.9c-0.5,3.4-1,6.7-1.5,10.1c-0.4,2.8,0.6,5.1,2.7,7    c5,4.8,10.5,7.7,17.7,7.9c11.1,0.3,22.2-0.4,33.2-1.4c9.7-0.9,19.3-1.8,29-1.8c9.9,0,19.1,2.9,28.4,5.6c10.3,3,19.8-3.9,18.9-13.5    c-0.5-4.9-3.8-8.7-8-9.1c-4.4-0.4-8.2,0.5-9.5,5.7c-3-4.1-1.3-7.2,2-9.5c6-4.2,8.2-3.2,16,6.4c0.5-4.7-4-10.6-10.4-12.5    c-3.7-1.1-7.6-1.5-11.4-2.1c-8-1.2-16.3-1-24.2-3c-9.4-2.4-15.3-8.8-19.1-17.3c-1.9-4.3,3.3-13.4,8.7-16.8    c5.2-3.2,11.1-1.8,16.7-2.7c4.1-0.7,6.7-1.3,6.3-6.2c4.3,4.9,4,6.3-1.5,7.6c4.7,0.3,9.3,1,13.9,2.2c4.8,1.3,10.5-3.6,12.2-9.5    c0.5,5.2-4.3,11.8-9.1,12.6c-5.5,1-10.7-2.2-16.8-1.9c9,5.1,9.8,5.2,16.8,3.2c-1.5,3.6-5.6,5.4-8.9,3.9c-3.8-1.7-7.5-3.7-11.4-5.4    c-5.9-2.6-11.8-3.5-17.1,1.9c4.5-1.8,8.8-3.9,13.9-1.9c4.7,1.8,8.2,5.6,12.9,7.6c-3.2,2.9-5.4,1.5-7-1.3c-2.2-3.8-5.1-5.8-10.4-5    c4.8,1.5,8.6,3.5,9,8.1c0.1,1.3-2.2,4.1-5.4,3.4c1.6-3.4,0.3-6-2.2-8.1c-2.7-2.2-6-2.4-9.2-1.1c-3.2,1.2-4.7,3.7-4.9,7.1    c-0.7,9.3,8,15,17,16.3c9.8,1.4,19.9,0.4,29.9,4.1c-0.6-4.1-2.2-8.4,1.2-11.5c2.9-2.5,6.6-3.4,10.6-2.5c-1.8,2.2-5.4,2.8-5.6,6.5    c-0.2,3.6,0.4,6.5,3.2,9.3c5.6,5.4,9.2,12,8.8,20.3c-0.2,4.9-2.6,8.9-4.8,13c-1.2,2.2-3,2-4.9,0.2c-0.1,2.5,1.9,2.1,2.3,3.2    c-1.6,1.4-2.7-1-4.1-0.5c-0.3,1.8,1.9,1.3,2.2,2.5c-2.2,1.1-4.3-0.5-6.4-0.6c0.6,2.7,3.3,1.1,4.5,2.5c-2.5,0.7-4.9,0.3-7.3-0.7    c0.2,2.4,2.1,1.8,3.2,2.4c-0.1,0.2-0.2,0.6-0.4,0.6c-1.3,0.4-2.6,1.9-4.1,0.7c-1.3-1.1-2.4-2.3-3.5-3.5c-3.5-4-7.9-5.5-13.3-5    c2.7,3,6.7,2.8,9.7,5.1c6.5,5.2,6.7,12.7,8,19.7c1.3,6.6,1,6.7,7.7,7.3c-3.2,2.7-6.5,4-10.5,0.5c0.1,1.1,0.1,1.8,0.3,1.9    c5.7,2.3,7.4,7.6,9.7,12.5c1.5,3.3,3.5,6.2,6.3,8.7c1,0.9,2.6,1.4,2,3.8c-0.5,2.4,2.3,1.4,3.2,1c2.5-1.1,5.1-2.1,6.9-4.8    c0.3,5.6-2.5,7.8-10.3,8.6c5.8,0.1,3.4,6.1,7.1,8c-3.4,1.4-4.6-2.3-7.1-2.2c-0.3,1.7,1.1,2.2,1.8,3c0.9,1.1,2.8,2,2,3.6    c-0.7,1.6-2.5-1.2-3.8,0.6c5.4,2.3,0.2,6.5,1.7,9.7c0.3,0.6-0.4,0.9-3.4-1.1c-0.8,1.4-0.6,2.9,0.8,3.4c4.2,1.4,3,3.6,1.2,6.1    c-2.1-0.9-2.8-4.4-5.6-3.5c-2.9,0.9-2.8,4-3.1,6.5c-2.4-0.2-3.4-2.4-5.5-2.9c1.2-2.4,5.9-0.4,5.1-4.2c-0.1-0.5-4,1.5-5.9,2.7    c-2.2,1.4-5.4-0.1-6.7,3.3c0.2-2.5,0.6-4.3-0.6-6.4c2,0,3.8-0.2,5.5,0c3.3,0.5,5.2-1.6,6.2-4c2.4-6,5.1-12.1,1.9-18.7    c-0.6-1.2-0.3-1.2,0.7-1.4c3.2-0.6,0.2-1.7,0-1.8c-4-0.4-6.1-3.6-8.7-6c-1.6-1.6-3.1-3.3-5.6-3.7c-2.8-0.4-3.4-3.1-4.8-5.7    c-0.4,3.3-2.1,5.5-3.8,7.9c0.6-2.9,0-6,1.8-8.5c1-1.4,0.6-2.5-0.5-2.7c-1.7-0.4,0,2-1.3,2.1c-9.2-7.2-15.5-16.7-21.6-26.4    c-2.6-4.2-5.7-8-13.3-6.8c6.7,2.5,10.1,7,14.5,11.2c-4,1.3-4.9-3.7-8-3.1c-0.2,1.7,1.2,1.9,2.3,2.6c2.5,1.5,1.6,3.6,0,4.8    c-1.5,1.2-1.2,2.6-0.9,3.5c0.4,1.2,2.3,0.8,3,0.3c2.7-1.7,4.9-0.8,7.2,0.7c-2.2,0.6-4.1,1.9-6.5,1.5c2.5,3.5,7.2,1.1,9.9,4.9    c-5.4-0.3-9.7-3.6-14.8-2.5c4.1,0.9,7.2,4.1,11.5,4.7c1.7,0.2,3.3,1.6,3.1,3.4c-0.6,6,2,10.5,6,14.5c0.6,0.6,0.9,0.7,0.3,1.3    c-0.5,0.6-1.1,0.6-1.5-0.1c-0.6-0.9-1.1-2.1-1.9-2.8c-1-0.8-2.2-0.7-3.1,0.4c-1.4,1.8,0.9,1,1.3,1.6c0.8,1.5,2.8,3.3,1.8,4.7    c-2.5,3.4-0.3,5.2,2.2,7.3c-5,1.6-8.2-0.3-10.7-4.2c-0.6,1.8,0,3.2,0.6,4.8c0.5,1.2,0.4,2.9-1,3.5c-1.6,0.6-1.6-1.5-2.6-2.6    c-0.8,4.2,3.4,4.9,4.5,7.9c-1.4,0.2-2.8,0.8-3.9-0.5c-1.4-1.8-2.3-1.6-4.2-0.3c-3,2-0.7-3.6-3.5-2.4c-0.6,1.3,1,2.4,1.1,3.8    c0.1,1.1,1.5,2.5,0.2,3.4c-1.3,0.9-1.8-1.5-3.2-1.4c-0.6,1.3-0.4,2.6,0.8,3.1c3.7,1.4,4.1,3.5,1.2,6.2c-0.5-0.4-1.3-1-1.9-1.8    c-2.9-3.5-5.3-2-5.9,1.3c-0.9,5.1-2.7,3-4.8,1.1c-0.3-0.3-0.8-0.4-1.3-0.7c1.2-2.2,5.3-0.6,5.7-5.1c-4.2,3.4-9.5,2.9-13.7,7.3    c0.5-2.8,1.2-4.8-0.3-7.1c8.2,2.4,11.7-3.3,14.8-8.7c2.2-3.8,4.3-7.5,7.2-10.9c3.1-3.6,2.1-8,1.3-12.2c-0.3-1.5-1-3-1.5-4.5    c2.1-0.2,3.4,2.2,5.6,1.3c-3.5-3.5-7.1-4.4-12.1-3c-3.1,0.9-7.6,1.7-9.5-3.5c4.5,4.2,8.1,3.2,11.5-1.2c0.9-1.2,2.6-1.7,4.3-2.7    c-1.6-0.8-3.9,0.4-4.6-2c0.7-0.3,1.6-0.6,2.4-0.9c-3.7,0.8-2.4-3.4-4.3-4.3c0.6-0.7,1.4-1.9,2.4-0.9c3.5,3.4,6,0.1,8.9-1.1    c-3.8-0.5-7.9,1.5-11.3-2.1c2.6-0.1,5.1-0.7,7.8-0.1c-1.5-1.2-3.1-2.4-5-3.9c0.9-0.1,1.6,0,2.2-0.3c1.5-0.6,3.5,2.1,4.5-0.7    c1-2.7-2.7-2.3-3.1-4.3c-0.3-1.5-1.1,0.1-1.6,0.3c-7.9,3.6-15.8,6.9-23.9,10c-3.3,1.3-4.5-2.7-7-1c0.3,1.2,1.2,1.8,2.7,1.8    c-0.6,0.5-0.7,0.8-1,0.8c-2,0.4-3.9-0.2-5.6-1.2c-0.8-0.4-1.6-1.5-2.3,0c-0.5,1.2,1.3,1,1.5,1.9c-2.4,0.8-4.3,0.1-6.3-0.9    c-3.2-1.7-2.6,0.2-1.2,2.6c-3.2-1-4.7-4-7.8-3.1c0.1,2.1,2.7,1.5,3.3,3.5c-4.6,0.6-8.9,4.2-13.6,0.3c2.9,0.4,5.8,2.1,8.4-0.6    c-8-1.6-9-2.7-7.8-8c1.9,0.1,2.7-1.3,3.6-2.7c-0.9,1.4-1.8,2.7-3.7,2.7c0.4-1.7,0.9-3.4,1.4-5.1c-2.2,0.2-1.9,2-2.1,3.6    c-1.2-1.2-1.5-3.5-3.2-2.7c-1.4,0.7,0,2.2,0.2,3.4c0.7,4.2,2.7,8.5,0,12.8c-3.7,0.1-3.8-9.4-10.4-3c4.9-1.6,6,2,7.5,4.4    c2,3-0.4,5.9-1.2,9.2c-0.8-1.8-1.4-3.3-1.9-4.7c-0.1,5.4-2.1,10.9,0.6,16c-3.2,0.2-4-0.6-5.7-6c-1.7,2.1-1.7,2.1-0.7,4.5    c1.2,2.9,2,6-0.2,8.7c-2.2,2.7-1.4-0.9-2.7-1.4c0,2.5,0,4.9,0,7.5c-2.1-1.3-0.2-3.5-1.8-4.4c-1.1,2.8,0.9,6.1-1,8.6    c-1.6,2.1-3.2,5.3-6.9,2.8c-0.6,2.5,1.2,2.9,2.6,3.6c-1.7,0.9-2.9,1.9-5.5,1.5c2.7,2.5,7.8,3.6,3,7.8c-2.2-1.6-4.1-3.7-5.6-6.8    c-0.7,1.2-1.6,2-1.7,2.9c-0.8,6.9-3.9,4-6.8,1.6c1.9-1.2,4.3-1.7,6-5c-5.1,2.4-9.2,4.3-13.7,6.4c-0.1-1.7,1.5-3.8-1.1-6    c10.7,1.8,13.8-5.1,16.6-12.3c3.3-8.3,6.1-16.8,8.5-25.5c0.9-3.3,0.4-6.6,0.4-9.9c0.1-8,0-8,6-10.2c1.4-1.2,2.7-0.4,3.8,0.3    C106.6,212.2,105.5,208.9,104.1,212.1z"
		pathData2 := "M87.1,216c-0.3,3.3,2,5.1,2.6,7.6c-1.5,0.9-2-0.4-2.6-1.3c-1.4-1.9-1.2-1.8-2.5,0c-1.2,1.6-2.9,0.2-4.2,0.6    c-1-2.2,2.3-2.6,1.5-4.5c-0.1,0-0.3-0.1-0.4-0.1c-1.9,0.3-2.4,4.6-4.6,2.6c-1.8-1.5,2.9-2.1,1.8-4.9c-2.2,1.4-3.8,5.1-6.9,2.8    c-2.1-1.6-3.8-3.6-6.1-5c-1.5,2.2,0.1,3.2,1.1,4.5c0.6,0.8,3.2,0.4,1.9,2.5c-1.1,1.7-3.6,1.9-4.3,0.8c-2.2-3.6-4.2,1.8-6-0.2    c-1.1,2,0.5,3.1,1.7,4.2c1.3,1.2,1.3,2.7,0.7,4.1c-0.4,0.9-1.7,0.7-2.2,0.2c-1.5-1.4-3.8-2.3-3.5-5c0.1-0.8,1.3-1.2,0.5-2.8    c-2.6,2.6-6.2,4.2-4.7,8.8c-0.6,0.1-0.9,0.2-0.9,0.1c-1.3-2.3-5.5-0.9-5.6-3.2c-0.1-1.4,3.1-2.9,4.8-4.4c3.6-3.2,3.7-8.1,5.6-12.1    c2.8-5.8,7.2-8.1,13.1-8.2c5.2-0.1,10.3,0,15.7,0c-0.9,1.8-3.8,2-3.4,5c2.4-1.9,4.2-4.4,7.9-3.9c-1.7,1.1-3,1.9-4.7,2.9    c3.1,2.2,4.7-0.9,6.8-1.1c0.5,0,1.3-1.7,1.6,0.1c0.1,0.7,0.4,1.2-1.3,1.4c-3.3,0.5-2.1,3.9-2.3,6.3c0,0.3,0.7,0.7,1,1    c0.7-1.2,1.4-2.3,2.1-3.5c1.8,0.7,1.8,2.3,1.4,3.7c-0.9,3.1-1.2,3.1-0.5,4.9c0.3,0.8,0.2,1.2-0.3,1.6c-1,0.9-1.5-0.2-1.8-0.7    C89.3,219.2,89.3,217.2,87.1,216z"
		pathData3 := "M123.9,138.7c0.4,5.5,6.3,7.8,4.8,13.1c-1.4-1.5-1.9-3.7-5.6-3.8c4.1,4.5,4.9,10.3,11.9,13    c-6-0.4-9.3-2.7-10.9-8.5c-1.9,8.9-2.5,11.6,5.5,18c-3.7,0.4-3.7,0.4-11-12.1c-0.9,4-0.3,6,2.9,10.3c-5.7-2.8-6.7-5.5-5.8-11.3    c1.6-10.8,3.2-21.7-2-32.2c-0.6-1.2-1.4-2.7,0.3-3.2c3-0.7,3.1-6.4,7.7-4.3c4.6,2.1,5.3,1.8,7.8-1.5c0.2,5.7-3.8,7.2-11.1,4.1    c1.7,3.3,1.7,3.4,7.3,3.8c-1.2,1.5-2.8,1.7-4.5,2.2c2.3,1.4,4,3.2,5.8,5.3c-2.4,1-3.1-3-6-1c5.3,2,4.2,11.5,12.4,9.8    C130.5,144.2,129.1,144.1,123.9,138.7z"
		pathData4 := "M106,101.7c-1.4,0.2-2.5,0.4-3.5,0.6c0.7,1.5,2.3,0.2,3.2,1.2c-1.1,0.5-2.3,1-3.4,1.5c-0.5,0.2-1.6,0.4-0.4,1    c0.9,0.4,2.3-0.3,1.2,1.8c-1.3,2.5-2.3,0.1-2.9-0.1c-1.9-1-5.5,0-5.6-2.3c-0.1-2.7,1.8-5.3,4-7.5c2-2,4.6-2.3,6.8-3.5    c0.7-0.4,1.1-0.1,1.6,0.5c1.5,1.7-3.1,0.1-1.2,2.3c1.4,1.7,0.4,2.9-1.8,3C103.6,100.1,104.2,100.9,106,101.7z"
		pathData5 := "M51.6,137.5c10.5-0.3,21-4.2,31.3,0.9C72.5,136,62,138.4,51.6,137.5z"
		pathData6 := "M80.2,133.5c-3.8,1.5-6.7-0.8-9.8-0.5c-1,0.1-1.4-0.5-1.4-1.5c0-1.4,0.9-1.5,2-1.5    C74.3,130,77.1,131.3,80.2,133.5z"
		pathData7 := "M83.1,108.1c0.1-3.5-1-6.8,1.7-10c0.1,3.3,0.6,6.3,3.2,8.5C86.4,107,85.1,107.5,83.1,108.1z"
		pathData8 := "M44.3,222.5c-1.4-7.2,2.6-5.1,5.5-4.8C48.1,219.1,46,219.9,44.3,222.5z"
		pathData9 := "M64.7,275c1.4-1.1,2.8-2.3,4.2-3.4c1.4-1.2,2.2-0.2,2.7,0.8c0.8,1.4-0.3,1.4-1.3,1.7c-1.6,0.5-3.2,1.3-4.7,2    C65.3,275.7,65,275.3,64.7,275z"
		pathData10 := "M214.9,274.5c2-1.5,3.3-2.4,4.6-3.5c0.8-0.6,1.4-1.5,2.4-0.5c0.4,0.4,1,1.1,0.9,1.6c0,1.3-1.2,0.7-1.9,0.9    C219.4,273.4,218.2,275.2,214.9,274.5z"
		pathData11 := "M155.9,274.7c1.5-1.6,2.8-3,4.2-4.4c1.1-1,1.9,0,2.4,0.7c0.8,1.1,0.5,1.9-1,1.9    C159.6,273,158.4,275.2,155.9,274.7z"
		pathData12 := "M79.5,108.8c-2.9-1.8-6-3.6-6.3-7.4C74.9,104.2,78,105.9,79.5,108.8z"
		pathData13 := "M65.7,140.2c3.6,0,7.1,0,10.7,0c0,0.2,0,0.3,0,0.5c-3.6,0-7.2,0-10.7,0C65.7,140.5,65.7,140.4,65.7,140.2z"
		pathData14 := "M45.8,230.5c1.9,1.9,1.1,3.5,1.1,5.1C45.1,233.8,45.1,233.8,45.8,230.5z"
		pathData15 := "M171.4,279.9c-1-2,0.4-3.2,1.2-4.6c0.4,0.3,0.9,0.6,1.1,0.9C174.5,278.5,172.4,278.8,171.4,279.9z"
		pathData16 := "M231.4,279.9c-1-2,0.4-3.2,1.2-4.6c0.4,0.3,0.9,0.6,1.1,0.9C234.5,278.5,232.4,278.8,231.4,279.9z"
		pathData17 := "M80.4,280.9c-1-2,0.4-3.2,1.2-4.6c0.4,0.3,0.9,0.6,1.1,1C83.5,279.5,81.4,279.8,80.4,280.9z"
		pathData18 := "M165.9,266.6c3-3.1,3.6-0.1,5.8,1.7C169.1,267.5,167.7,267.1,165.9,266.6z"
		pathData19 := "M230.7,267.5c-1.8,0.7-2.6-1.7-4.3-0.6C228.6,263.3,229.4,266.9,230.7,267.5z"
		pathData20 := "M74.7,267.5c2.9-1.2,2.9-1.2,5.1,1.5C78.1,268.5,76.7,268.1,74.7,267.5z"
		pathData21 := "M249.4,277.2c-0.5-2.1-1-3.7,1-4.9C251.6,274,251.1,275.3,249.4,277.2z"
		pathData22 := "M99.8,275.1c0.3,2.2,0.6,3.8-1.5,4.7C97.8,278,97.5,276.3,99.8,275.1z"
		pathData23 := "M189.4,277.2c-0.5-2.1-1-3.7,1-4.9C191.6,274,191.1,275.3,189.4,277.2z"
		pathData24 := "M67.1,133c-2.1-1.1-2.5-2.2-1.8-3.6c0.1-0.2,1.2-0.2,1.2-0.2C66.7,130.3,66.9,131.4,67.1,133z"
		pathData25 := "M40.5,224.3c-0.8-2.4-0.3-3.5,1.1-4.1C42.7,221.8,40.7,222.5,40.5,224.3z"
		pathData26 := "M68.7,226.7c-1-1.3-0.8-2.3,0.3-3.1c0.5-0.4,0.9-0.4,1,0.1C70.3,225,68.3,225.3,68.7,226.7z"
		pathData27 := "M60,126.5c0.9,0.7,2.1,1.4,1.4,2.5c-0.2,0.3-0.8,0.5-1.2,0.7C60.2,128.6,60.1,127.6,60,126.5z"
		pathData28 := "M76.5,116.9c9-5.4,13.6,0.5,17.7,7.3C88.7,120.8,80.8,123,76.5,116.9z"
		pathData29 := "M187.5,200.1c-11.4-3.8-21.3,3-31.9,4c-0.9-2.2,4.4,0.5,1.9-2.9c-0.1-0.2,0.7-1,1.5-0.7    c0.4,0.2,0.9,0.5,1.3,0.5c3.7-0.4,7.3-1.4,10.7-2.4C177.5,196.8,181.8,196.2,187.5,200.1z"
		pathData30 := "M107.3,116.8c7-10.5,7-10.5,13.9-13.1c-2.3,4.5-6.9,6-7.8,10.2c-0.4-0.1,0.4,0.1,1.1,0.3    C112.7,116.4,110.5,117.6,107.3,116.8z"
		pathData31 := "M105.3,219.9c1.5-1.1,2.7-1.9,2.9-3.6c2.7,1.9-1.2,4,0.4,6.1c0.2,0.3-1.7,0.9-2.5,0.1c-2.5-2.3-4.3-5-5.7-8.2    C103,215,104.9,216.5,105.3,219.9z"
		pathData32 := "M200.8,226.7c-4.3-4.4-10.8-1.4-15.2-5C188.9,219.5,198.5,222.5,200.8,226.7z"
		pathData33 := "M222.7,193.7c-1.6-4.6-5.3-6.9-8.7-9.2c3.9-0.7,10.2,3.3,11,7.5c0.4,2-0.8,2.3-2.4,1.6L222.7,193.7z"
		pathData34 := "M89.6,127.4c-4.2,1.3,0.6,2.1,0.1,3.6c-5.4-1.4-7.7-6.4-11.3-9.7C82.5,122.6,85.1,126.8,89.6,127.4z"
		pathData35 := "M96.6,184c2.1-4.6,7-6.2,10-9.7C105.4,179.8,100.7,181.6,96.6,184z"
		pathData36 := "M84.2,143.5c3.8-2.4,4.8-5.6,3.3-9.6c-0.3-0.9,0.3-1.3,0.7-1.4c1.4-0.3,2.1,1,2.3,1.8    C91.7,138.2,88.9,142.4,84.2,143.5z"
		pathData37 := "M203.7,246.6c-5.1-2.1-6.8-6.8-11.1-7.8c-0.1,0-0.2-0.6-0.1-0.7c0.3-0.4,1-1,1.1-0.9    C197.2,239.5,200.9,241.6,203.7,246.6z"
		pathData38 := "M198.3,238.8c-2.2-2.6-3.9-4.6-5.6-6.6c-0.4-0.4-2,0.2-1.1-1.3c0.4-0.7,1.3-1.5,2-0.7    C195.6,232.6,199.3,233.8,198.3,238.8z"
		pathData39 := "M157.6,209.2c-4.4,1.2-8.9,3.3-13.4,0.1C148.7,212.2,153,207.7,157.6,209.2z"
		pathData40 := "M222.6,193.6c2.2,2.4,3.8,5.1,4.8,8c-3.2-0.2-6.3-5-4.7-7.9C222.7,193.7,222.6,193.6,222.6,193.6z"
		pathData41 := "M193.3,226.7c3.5,0.1,6.5,0.9,7.4,5.1c-4.1,0.7-4.7-3.8-7.3-5.2L193.3,226.7z"
		pathData42 := "M175.4,204.3c-0.3-1.7-2.7-1.1-3-3.1c2.9-1.1,5.3,1.7,8.5,0.5c-1.6,2.5-3.8,2.1-5.7,2.4L175.4,204.3z"
		pathData43 := "M123.6,209.5c2.9-1.1,2.6-4.7,5.3-6.5c0.4,4.3-0.9,6.6-5.2,6.4C123.7,209.4,123.6,209.5,123.6,209.5z"
		pathData44 := "M104.1,212.1c-0.2-1.3-0.3-2.7-0.5-4.3c2.8,1.1,3.5,5.1,8,4c-3.4,2.5-5.4-0.5-7.6,0.2    C103.9,211.9,104.1,212.1,104.1,212.1z"
		pathData45 := "M175.2,204.2c-2.8,1.3-5.7,3.1-8.9,0.1c3.2,0,6.1,0,9,0C175.4,204.3,175.2,204.2,175.2,204.2z"
		pathData46 := "M183.2,188.3c4.5-0.6,6.7,1.1,8.6,3.4C189.3,190.7,186.8,189.7,183.2,188.3z"
		pathData47 := "M87.9,113.1c-2.1,0.3-4.3,0.6-6.4,1c2-2.5,4.1-3,6.3-0.8L87.9,113.1z"
		pathData48 := "M159.2,208.8c1.8-3.8,4.3-1,6.3-1.4C163.5,208.1,161.7,209.6,159.2,208.8z"
		pathData49 := "M193.1,242.4c0.5-0.2,0.6-0.4,0.7-0.4c2,0.8,3.3,2.4,4.7,3.9c-0.2,0.2-0.6,0.7-0.7,0.6    C195.7,245.9,194.6,243.9,193.1,242.4z"
		pathData50 := "M87.4,165.1c-1.1-3.3,1.3-3.9,3.1-6C90.6,162.3,89,163.5,87.4,165.1z"
		pathData51 := "M225.8,205.8c2.8,0.2,5,0.9,2.8,3.9C227.6,208.4,226.7,207.1,225.8,205.8z"
		pathData52 := "M87.7,113.3c2.3-0.5,3.9,0.7,5.2,2.4c0.1,0.2-0.3,1.2-0.4,1.2c-2.3-0.3-3.3-2.2-4.6-3.7    C87.9,113.1,87.7,113.3,87.7,113.3z"
		pathData53 := "M195.1,191c1,0.3,1.8,1,1.9,2.3c0,0.8-0.1,2-1.1,1.5c-0.8-0.4-1.4-1.5-1.8-2.5    C193.8,191.8,194.1,191.1,195.1,191z"
		pathData54 := "M175.7,189.9c1.6,0.8,3.3,1.6,5.1,2.5C177.5,192.9,177.5,192.9,175.7,189.9z"
		pathData55 := "M128.2,209.6c1.7-2.8,4.1-0.5,5.7-2C132.1,208.6,130.7,210.8,128.2,209.6z"
		pathData56 := "M189.2,254c3.4,2.2-0.1,3.8,0,5.6C189.2,257.9,189.2,256.2,189.2,254z"
		pathData57 := "M193.4,226.6c-2.3,0.2-4.6,0.3-7,0.5c2.3-1,4.6-1.7,6.8-0.3C193.3,226.7,193.4,226.6,193.4,226.6z"
		pathData58 := "M189.5,240.2c-0.7-2.3-0.9-3.9,1.5-5.2C190.5,236.8,190.1,238.2,189.5,240.2z"
		pathData59 := "M83.9,174.9c1-1.2,1.7-2.7,4.4-3c-1.6,1.7-2.2,3.6-4.6,2.8C83.7,174.8,83.9,174.9,83.9,174.9z"
		pathData60 := "M199.7,197.1c-0.7-0.8-1.2-1.2-1.5-1.7c-0.3-0.6-0.1-1.2,0.7-1.4c0.8-0.1,1.2,0.3,1.1,1.1    C199.9,195.7,199.8,196.2,199.7,197.1z"
		pathData61 := "M159.7,193.6c2.2-0.8,3.6-0.8,5-0.1C163.3,194.3,162,194.1,159.7,193.6z"
		pathData62 := "M171,192.2c-0.3,0.7-0.8,1-1.4,0.6c-0.5-0.3-0.8-0.9-1.3-1.3c0.6-0.1,1.1-0.4,1.7-0.4    C170.8,191,171,191.5,171,192.2z"
		pathData63 := "M79.5,209.2c1.7,5-2.7,5.4-5.6,8.2C74.5,212.9,79.9,213.3,79.5,209.2z"
		pathData64 := "M86.7,211.6c0.5,3.6-1.2,5.4-3.5,7C83.9,216.2,84.7,213.9,86.7,211.6z"
		pathData65 := "M75.7,211.5c-0.8,1.3-1.6,1.6-2.7,1.4c-1.8-0.4-0.2-1-0.1-1.5c0.1-1.9,0.9-0.9,1.7-0.5    C75.1,211.2,75.9,211,75.7,211.5z"
		pathData66 := "M82.4,207.6c1,2.4,0.5,3.5-1.1,4.2C80.5,210.2,82.2,209.6,82.4,207.6z"
		pathData67 := "M83.6,120c-2.7-2.5-1.1-3.3,1.2-4.5C84.3,117.3,84,118.6,83.6,120z"
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
		canvas.Path(pathData20, "fill:"+hexCode)
		canvas.Path(pathData21, "fill:"+hexCode)
		canvas.Path(pathData22, "fill:"+hexCode)
		canvas.Path(pathData23, "fill:"+hexCode)
		canvas.Path(pathData24, "fill:"+hexCode)
		canvas.Path(pathData25, "fill:"+hexCode)
		canvas.Path(pathData26, "fill:"+hexCode)
		canvas.Path(pathData27, "fill:"+hexCode)
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
		canvas.Path(pathData67, "fill:"+hexCode)
	},
}
