/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/boomsquared/tap/pkg/file"
	"github.com/spf13/cobra"
)

var by string

// groupCmd represents the group command
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Automatically group files to directory",
	Long: `group -by flags
		ImageWidth
		ImageLength
		BitsPerSample
		Compression
		PhotometricInterpretation
		Orientation
		SamplesPerPixel
		PlanarConfiguration
		YCbCrSubSampling
		YCbCrPositioning
		XResolution
		YResolution
		ResolutionUnit
		DateTime
		ImageDescription
		Make
		Model
		Software
		Artist
		Copyright
		ExifIFDPointer
		GPSInfoIFDPointer
		InteroperabilityIFDPointer
		ExifVersion
		FlashpixVersion
		ColorSpace
		ComponentsConfiguration
		CompressedBitsPerPixel
		PixelXDimension
		PixelYDimension
		MakerNote
		UserComment
		RelatedSoundFile
		DateTimeOriginal
		DateTimeDigitized
		SubSecTime
		SubSecTimeOriginal
		SubSecTimeDigitized
		ImageUniqueID
		ExposureTime
		FNumber
		ExposureProgram
		SpectralSensitivity
		ISOSpeedRatings
		OECF
		ShutterSpeedValue
		ApertureValue
		BrightnessValue
		ExposureBiasValue
		MaxApertureValue
		SubjectDistance
		MeteringMode
		LightSource
		Flash
		FocalLength
		SubjectArea
		FlashEnergy
		SpatialFrequencyResponse
		FocalPlaneXResolution
		FocalPlaneYResolution
		FocalPlaneResolutionUnit
		SubjectLocation
		ExposureIndex
		SensingMethod
		FileSource
		SceneType
		CFAPattern
		CustomRendered
		ExposureMode
		WhiteBalance
		DigitalZoomRatio
		FocalLengthIn35mmFilm
		SceneCaptureType
		GainControl
		Contrast
		Saturation
		Sharpness
		DeviceSettingDescription
		SubjectDistanceRange
		LensMake
		LensModel`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a path argument")
		}
		if _, err := os.Stat(args[0]); os.IsNotExist(err) {
			return errors.New("path does not exist")
		}

		fs, err := os.Stat(args[0])
		if err != nil {
			return err
		}
		if !fs.IsDir() {
			return errors.New("path is not a directory")
		}

		return nil

	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("BY: ", by)
		an := file.NewAssigner(file.Operation{})
		if err := an.Load(args[0], by); err != nil {
			fmt.Println(err)
		}
		if err := an.Assign(); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(groupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// groupCmd.PersistentFlags().String("foo", "", "A help for foo")

	groupCmd.Flags().StringVar(&by, "by", "model", "file will be assigned to folder base on this flag")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// groupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
