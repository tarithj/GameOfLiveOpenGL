package main

const (
	vertexShaderSource = `
    #version 410
    in vec3 vp;
    out vec4 frag_colour;

    void main() {
        gl_Position = vec4(vp, 1.0);
    }
` + "\x00"

	fragmentShaderSource = `
    #version 410
    out vec4 frag_colour;
    void main() {
		// color
        frag_colour = vec4(18, 18, 18, 1);
    }
` + "\x00"
)
