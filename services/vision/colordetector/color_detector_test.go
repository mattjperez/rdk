package colordetector

import (
	"context"
	"testing"

	"go.viam.com/test"
	"go.viam.com/utils/artifact"

	"go.viam.com/rdk/components/camera"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/rimage"
	"go.viam.com/rdk/services/vision"
	"go.viam.com/rdk/testutils/inject"
	"go.viam.com/rdk/vision/objectdetection"
)

func TestColorDetector(t *testing.T) {
	inp := objectdetection.ColorDetectorConfig{
		SegmentSize:       150000,
		HueTolerance:      0.44,
		DetectColorString: "#4F3815",
	}
	ctx := context.Background()
	r := &inject.Robot{}
	r.LoggerFunc = func() logging.Logger { return nil }
	name := vision.Named("test_cd")
	srv, err := registerColorDetector(ctx, name, &inp, r)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, srv.Name(), test.ShouldResemble, name)
	img, err := rimage.NewImageFromFile(artifact.MustPath("vision/objectdetection/detection_test.jpg"))
	test.That(t, err, test.ShouldBeNil)

	// Test properties. Should support detections and not classifications or object PCDs
	props, err := srv.GetProperties(ctx, nil)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, props.DetectionSupported, test.ShouldEqual, true)
	test.That(t, props.ClassificationSupported, test.ShouldEqual, false)
	test.That(t, props.ObjectPCDsSupported, test.ShouldEqual, false)

	// Does implement Detections
	det, err := srv.Detections(ctx, img, nil)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, det, test.ShouldHaveLength, 1)

	// Does not implement Classifications
	_, err = srv.Classifications(ctx, img, 1, nil)
	test.That(t, err, test.ShouldNotBeNil)
	test.That(t, err.Error(), test.ShouldContainSubstring, "does not implement")

	// with error - bad parameters
	inp.HueTolerance = 4.0 // value out of range
	_, err = registerColorDetector(ctx, name, &inp, r)
	test.That(t, err.Error(), test.ShouldContainSubstring, "hue_tolerance_pct is required")

	// with error - nil parameters
	_, err = registerColorDetector(ctx, name, nil, r)
	test.That(t, err.Error(), test.ShouldContainSubstring, "cannot be nil")
}

func TestRegistrationWithDefaultCamera(t *testing.T) {
	ctx := context.Background()
	serviceName := vision.Named("color_detector")
	cameraName := camera.Named("test")
	modelCfg := objectdetection.ColorDetectorConfig{
		SegmentSize:       150000,
		HueTolerance:      0.44,
		DetectColorString: "#4F3815",
	}

	r := &inject.Robot{}
	r.LoggerFunc = func() logging.Logger { return nil }
	r.ResourceByNameFunc = func(name resource.Name) (resource.Resource, error) {
		if name == cameraName {
			return inject.NewCamera(cameraName.Name), nil
		}
		return nil, resource.NewNotFoundError(name)
	}

	service, err := registerColorDetector(ctx, serviceName, &modelCfg, r)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, service, test.ShouldNotBeNil)

	modelCfg.DefaultCamera = cameraName.Name
	service, err = registerColorDetector(ctx, serviceName, &modelCfg, r)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, service, test.ShouldNotBeNil)

	modelCfg.DefaultCamera = "not-camera"
	_, err = registerColorDetector(ctx, serviceName, &modelCfg, r)
	test.That(t, err, test.ShouldNotBeNil)
	test.That(t, err.Error(), test.ShouldContainSubstring, "could not find camera \"not-camera\"")
}
