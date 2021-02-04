package main

const (
	vertexShaderSource = `
    #version 410
    in vec3 vp;

    void main() {
        gl_Position = vec4(vp, 1.0);
    }
` + "\x00"

	fragmentShaderSource = `
    #version 410
    out vec4 frag_colour;
    void main() {
		// color
        frag_colour = vec4(0.255, 0.179, 0.67,1);
    }
` + "\x00"
)
