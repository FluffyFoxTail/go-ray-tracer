At its core, a ray tracer sends rays through pixels and computes the COLOR seen 
in the direction of those rays. 
The involved steps are:

* Calculate the ray from the “eye” through the pixel,
* Determine which objects the ray intersects, and
* Compute a COLOR for the closest intersection point.