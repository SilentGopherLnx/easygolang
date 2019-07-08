package easygtk

import (
	. "github.com/SilentGopherLnx/easygolang"

	"bytes"
	"image"
	"image/png"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	//"github.com/gotk3/gotk3/cairo"
	//"github.com/gotk3/gotk3/glib"
)

func GTK_PixBuf_From_Bytes(data *[]byte, ftype string) *gdk.Pixbuf {
	if data != nil && len(*data) > 0 && ftype != "" {
		loader, err1 := gdk.PixbufLoaderNewWithType(ftype)
		if err1 == nil {
			pixbuf, err2 := loader.WriteAndReturnPixbuf(*data)
			if err2 == nil {
				return pixbuf
			}
		}
	}
	return nil
}

func GTK_PixBuf_From_RGBA(img image.Image) *gdk.Pixbuf {
	if InterfaceNil(img) {
		return nil
	}
	data := []byte{}
	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	if err == nil {
		data = buf.Bytes()
		return GTK_PixBuf_From_Bytes(&data, "png")
	} else {
		return nil
	}

	/*	w := img.Bounds().Max.X
		h := img.Bounds().Max.Y

		pixbuf, _ := gdk.PixbufNew(gdk.COLORSPACE_RGB, true, 8, w, h)

		//Prln("wh" + I2S(w) + "/" + I2S(h) + " " + TypeOf(pixbuf) + "/" + B2S_YN(pixbuf == nil))

		surf := cairo.CreateImageSurface(cairo.FORMAT_ARGB32, w, h)
		context := cairo.Create(surf)
		gtk.GdkCairoSetSourcePixBuf(context, pixbuf, 0, 0)
		context.Paint()
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				col := img.RGBAAt(x, y)
				context.SetSourceRGBA(float64(col.R)/255.0, float64(col.G)/255.0, float64(col.B)/255.0, float64(col.A)/255.0)
				context.Rectangle(float64(x), float64(y), 1, 1)
				context.Fill()
			}
		}
		surf.Flush()

		context.Close()
		surf.Close()*/

	//loader, _ := gdk.PixbufLoaderNew()
	//loader.WriteAndReturnPixbuf()
	//pixbuf, _ = pixbuf.ScaleSimple(w, h, gdk.INTERP_NEAREST)
	//surf.WriteToPNG("/mnt/dm-1/golang/my_code/FileManager/mime/out.png")

	//return pixbuf

	/* https://gist.github.com/bert/985903/1f7675464104c70fb65ed93328b7538f459775dc

		    g_type_init ();
	        pixbuf = gdk_pixbuf_new_from_file ("test.png", NULL);
	        g_assert (pixbuf != NULL);
	        format = (gdk_pixbuf_get_has_alpha (pixbuf)) ? CAIRO_FORMAT_ARGB32 : CAIRO_FORMAT_RGB24;
	        width = gdk_pixbuf_get_width (pixbuf);
	        height = gdk_pixbuf_get_height (pixbuf);
	        surface = cairo_image_surface_create (format, width, height);
	        g_assert (surface != NULL);
	        cr = cairo_create (surface);

	        gdk_cairo_set_source_pixbuf (cr, pixbuf, 0, 0);
	        cairo_paint (cr);

	        cairo_set_source_rgb (cr, 1, 0, 0);
	        cairo_rectangle (cr, width * .25, height * .25, width *.5, height *.5);
	        cairo_fill (cr);

	        cairo_surface_write_to_png (surface, "output.png");
	        cairo_surface_destroy (surface);
	        cairo_destroy (cr);*/
}

// ===========

func GTK_Image_From_PixBuf(pixbuf *gdk.Pixbuf) *gtk.Image {
	im, err := gtk.ImageNew()
	if err == nil {
		if pixbuf != nil {
			im.SetFromPixbuf(pixbuf)
		}
		return im
	}
	return nil
}

func GTK_Image_From_Name(name string, size gtk.IconSize) *gtk.Image {
	im, err := gtk.ImageNew()
	if err == nil {
		im.SetFromIconName(name, size)
		return im
	}
	return nil
}

func GTK_Image_From_File(fname string, ftype string) *gtk.Image {
	//pixbuf, err := gdk.PixbufNewFromFile(fname)
	data, ok := FileBytesRead(fname)
	if ok {
		pixbuf := GTK_PixBuf_From_Bytes(data, ftype)
		return GTK_Image_From_PixBuf(pixbuf)
	}
	im, _ := gtk.ImageNew()
	return im
}

//img, _, err1 := image.Decode(bytes.NewReader(*data))
//if err != nil {fmt.Println(err1)}
//wh := img.Bounds().Max
