#ifndef PHYSICSGEOM_H
#define PHYSICSGEOM_H


#ifdef __cplusplus
extern "C" {
#endif

typedef struct { float x, y, z; } Vector;
typedef struct { Vector c1, c2, c3; } Matrix;
typedef struct { float x, y, z, w; } Quaternion;
typedef struct { Vector p; Quaternion o; } Transform;

Quaternion Matrix_toQuat(Matrix m);

#ifdef __cplusplus
}
#endif


#endif // PHYSICSGEOM_H

