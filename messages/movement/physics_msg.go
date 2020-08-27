package movement

import "fmt"

/*
Structure definition for Physics Messages (protocol 16)
PhysicsBehavior Messages
*/

/*
Physics Behavior
*/
type PhysicsState struct {
	PositionX        float32
	PositionY        float32
	PositionZ        float32
	RotationX        float32
	RotationY        float32
	RotationZ        float32
	VelocityX        float32
	VelocityY        float32
	VelocityZ        float32
	AngularVelocityX float32
	AngularVelocityY float32
	AngularVelocityZ float32
	GlobalID         uint64
}

func (m PhysicsState) String() string {
	return fmt.Sprintf("%T", m)
}

/*
Physics Behavior
*/
type PhysicsForce struct {
	PositionX        float32
	PositionY        float32
	PositionZ        float32
	RotationX        float32
	RotationY        float32
	RotationZ        float32
	VelocityX        float32
	VelocityY        float32
	VelocityZ        float32
	AngularVelocityX float32
	AngularVelocityY float32
	AngularVelocityZ float32
	ForceX           float32
	ForceY           float32
	ForceZ           float32
	GlobalID         uint64
}

func (m PhysicsForce) String() string {
	return fmt.Sprintf("%T", m)
}

/*
Physics Behavior
*/
type PhysicsForceAtPosition struct {
	PositionX        float32
	PositionY        float32
	PositionZ        float32
	RotationX        float32
	RotationY        float32
	RotationZ        float32
	VelocityX        float32
	VelocityY        float32
	VelocityZ        float32
	AngularVelocityX float32
	AngularVelocityY float32
	AngularVelocityZ float32
	ForceX           float32
	ForceY           float32
	ForceZ           float32
	ForcePosX        float32
	ForcePosY        float32
	ForcePosZ        float32
	GlobalID         uint64
}

func (m PhysicsForceAtPosition) String() string {
	return fmt.Sprintf("%T", m)
}

/*
Physics Behavior
*/
type PhysicsTorque struct {
	PositionX        float32
	PositionY        float32
	PositionZ        float32
	RotationX        float32
	RotationY        float32
	RotationZ        float32
	VelocityX        float32
	VelocityY        float32
	VelocityZ        float32
	AngularVelocityX float32
	AngularVelocityY float32
	AngularVelocityZ float32
	TorqueX          float32
	TorqueY          float32
	TorqueZ          float32
	GlobalID         uint64
}

func (m PhysicsTorque) String() string {
	return fmt.Sprintf("%T", m)
}

type PhysicsGrab struct {
	GlobalID uint64
	HolderID uint64
	OffsetX  float32
	OffsetY  float32
	OffsetZ  float32
	Force    float32
}

func (m PhysicsGrab) String() string {
	return fmt.Sprintf("%T", m)
}

type PhysicsRelease struct {
	GlobalID uint64
	HolderID uint64
}

func (m PhysicsRelease) String() string {
	return fmt.Sprintf("%T", m)
}
