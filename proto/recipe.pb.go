// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1
// 	protoc        v3.11.4
// source: recipe.proto

package mealz

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Ethnicity int32

const (
	Ethnicity_UNKNOWN_ETHNICITY Ethnicity = 0
	Ethnicity_AMERICAN          Ethnicity = 1
	Ethnicity_CARIBBEAN         Ethnicity = 2
	Ethnicity_CHINESE           Ethnicity = 3
	Ethnicity_ASIAN             Ethnicity = 4
	Ethnicity_ITALIAN           Ethnicity = 5
	Ethnicity_THAI              Ethnicity = 6
	Ethnicity_CAJUN             Ethnicity = 7
	Ethnicity_JAPANESE          Ethnicity = 8
	Ethnicity_MOROCCAN          Ethnicity = 9
)

// Enum value maps for Ethnicity.
var (
	Ethnicity_name = map[int32]string{
		0: "UNKNOWN_ETHNICITY",
		1: "AMERICAN",
		2: "CARIBBEAN",
		3: "CHINESE",
		4: "ASIAN",
		5: "ITALIAN",
		6: "THAI",
		7: "CAJUN",
		8: "JAPANESE",
		9: "MOROCCAN",
	}
	Ethnicity_value = map[string]int32{
		"UNKNOWN_ETHNICITY": 0,
		"AMERICAN":          1,
		"CARIBBEAN":         2,
		"CHINESE":           3,
		"ASIAN":             4,
		"ITALIAN":           5,
		"THAI":              6,
		"CAJUN":             7,
		"JAPANESE":          8,
		"MOROCCAN":          9,
	}
)

func (x Ethnicity) Enum() *Ethnicity {
	p := new(Ethnicity)
	*p = x
	return p
}

func (x Ethnicity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Ethnicity) Descriptor() protoreflect.EnumDescriptor {
	return file_recipe_proto_enumTypes[0].Descriptor()
}

func (Ethnicity) Type() protoreflect.EnumType {
	return &file_recipe_proto_enumTypes[0]
}

func (x Ethnicity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Ethnicity.Descriptor instead.
func (Ethnicity) EnumDescriptor() ([]byte, []int) {
	return file_recipe_proto_rawDescGZIP(), []int{0}
}

type Season int32

const (
	Season_UNKNOWN_SEASON Season = 0
	Season_WINTER         Season = 1
	Season_SPRING         Season = 2
	Season_SUMMER         Season = 3
	Season_FALL           Season = 4
)

// Enum value maps for Season.
var (
	Season_name = map[int32]string{
		0: "UNKNOWN_SEASON",
		1: "WINTER",
		2: "SPRING",
		3: "SUMMER",
		4: "FALL",
	}
	Season_value = map[string]int32{
		"UNKNOWN_SEASON": 0,
		"WINTER":         1,
		"SPRING":         2,
		"SUMMER":         3,
		"FALL":           4,
	}
)

func (x Season) Enum() *Season {
	p := new(Season)
	*p = x
	return p
}

func (x Season) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Season) Descriptor() protoreflect.EnumDescriptor {
	return file_recipe_proto_enumTypes[1].Descriptor()
}

func (Season) Type() protoreflect.EnumType {
	return &file_recipe_proto_enumTypes[1]
}

func (x Season) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Season.Descriptor instead.
func (Season) EnumDescriptor() ([]byte, []int) {
	return file_recipe_proto_rawDescGZIP(), []int{1}
}

type UnitOfMeasure int32

const (
	UnitOfMeasure_UNKNOWN_UNITOFMEASURE UnitOfMeasure = 0
	UnitOfMeasure_TEASPOON              UnitOfMeasure = 1
	UnitOfMeasure_TABLESPOON            UnitOfMeasure = 2
	UnitOfMeasure_POUND                 UnitOfMeasure = 3
	UnitOfMeasure_OUNCE                 UnitOfMeasure = 4
	UnitOfMeasure_CUP                   UnitOfMeasure = 5
	UnitOfMeasure_PIECE                 UnitOfMeasure = 6
)

// Enum value maps for UnitOfMeasure.
var (
	UnitOfMeasure_name = map[int32]string{
		0: "UNKNOWN_UNITOFMEASURE",
		1: "TEASPOON",
		2: "TABLESPOON",
		3: "POUND",
		4: "OUNCE",
		5: "CUP",
		6: "PIECE",
	}
	UnitOfMeasure_value = map[string]int32{
		"UNKNOWN_UNITOFMEASURE": 0,
		"TEASPOON":              1,
		"TABLESPOON":            2,
		"POUND":                 3,
		"OUNCE":                 4,
		"CUP":                   5,
		"PIECE":                 6,
	}
)

func (x UnitOfMeasure) Enum() *UnitOfMeasure {
	p := new(UnitOfMeasure)
	*p = x
	return p
}

func (x UnitOfMeasure) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UnitOfMeasure) Descriptor() protoreflect.EnumDescriptor {
	return file_recipe_proto_enumTypes[2].Descriptor()
}

func (UnitOfMeasure) Type() protoreflect.EnumType {
	return &file_recipe_proto_enumTypes[2]
}

func (x UnitOfMeasure) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UnitOfMeasure.Descriptor instead.
func (UnitOfMeasure) EnumDescriptor() ([]byte, []int) {
	return file_recipe_proto_rawDescGZIP(), []int{2}
}

type Ingredient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId      []byte        `protobuf:"bytes,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	Title         string        `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	UnitOfMeasure UnitOfMeasure `protobuf:"varint,3,opt,name=unit_of_measure,json=unitOfMeasure,proto3,enum=mealz.UnitOfMeasure" json:"unit_of_measure,omitempty"`
	Amount        float32       `protobuf:"fixed32,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Vegetarian    bool          `protobuf:"varint,5,opt,name=vegetarian,proto3" json:"vegetarian,omitempty"`
}

func (x *Ingredient) Reset() {
	*x = Ingredient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recipe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ingredient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ingredient) ProtoMessage() {}

func (x *Ingredient) ProtoReflect() protoreflect.Message {
	mi := &file_recipe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ingredient.ProtoReflect.Descriptor instead.
func (*Ingredient) Descriptor() ([]byte, []int) {
	return file_recipe_proto_rawDescGZIP(), []int{0}
}

func (x *Ingredient) GetObjectId() []byte {
	if x != nil {
		return x.ObjectId
	}
	return nil
}

func (x *Ingredient) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Ingredient) GetUnitOfMeasure() UnitOfMeasure {
	if x != nil {
		return x.UnitOfMeasure
	}
	return UnitOfMeasure_UNKNOWN_UNITOFMEASURE
}

func (x *Ingredient) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Ingredient) GetVegetarian() bool {
	if x != nil {
		return x.Vegetarian
	}
	return false
}

type Recipe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId    []byte        `protobuf:"bytes,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	Title       string        `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Vegetarian  bool          `protobuf:"varint,3,opt,name=vegetarian,proto3" json:"vegetarian,omitempty"`
	Ethnicity   Ethnicity     `protobuf:"varint,4,opt,name=ethnicity,proto3,enum=mealz.Ethnicity" json:"ethnicity,omitempty"`
	Season      Season        `protobuf:"varint,5,opt,name=season,proto3,enum=mealz.Season" json:"season,omitempty"`
	Source      string        `protobuf:"bytes,6,opt,name=source,proto3" json:"source,omitempty"`
	Ingredients []*Ingredient `protobuf:"bytes,7,rep,name=ingredients,proto3" json:"ingredients,omitempty"`
}

func (x *Recipe) Reset() {
	*x = Recipe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recipe_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recipe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recipe) ProtoMessage() {}

func (x *Recipe) ProtoReflect() protoreflect.Message {
	mi := &file_recipe_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recipe.ProtoReflect.Descriptor instead.
func (*Recipe) Descriptor() ([]byte, []int) {
	return file_recipe_proto_rawDescGZIP(), []int{1}
}

func (x *Recipe) GetObjectId() []byte {
	if x != nil {
		return x.ObjectId
	}
	return nil
}

func (x *Recipe) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Recipe) GetVegetarian() bool {
	if x != nil {
		return x.Vegetarian
	}
	return false
}

func (x *Recipe) GetEthnicity() Ethnicity {
	if x != nil {
		return x.Ethnicity
	}
	return Ethnicity_UNKNOWN_ETHNICITY
}

func (x *Recipe) GetSeason() Season {
	if x != nil {
		return x.Season
	}
	return Season_UNKNOWN_SEASON
}

func (x *Recipe) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Recipe) GetIngredients() []*Ingredient {
	if x != nil {
		return x.Ingredients
	}
	return nil
}

type RecipeGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId []byte `protobuf:"bytes,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
}

func (x *RecipeGetRequest) Reset() {
	*x = RecipeGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recipe_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecipeGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecipeGetRequest) ProtoMessage() {}

func (x *RecipeGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recipe_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecipeGetRequest.ProtoReflect.Descriptor instead.
func (*RecipeGetRequest) Descriptor() ([]byte, []int) {
	return file_recipe_proto_rawDescGZIP(), []int{2}
}

func (x *RecipeGetRequest) GetObjectId() []byte {
	if x != nil {
		return x.ObjectId
	}
	return nil
}

type RecipeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Recipe *Recipe `protobuf:"bytes,1,opt,name=recipe,proto3" json:"recipe,omitempty"`
}

func (x *RecipeRequest) Reset() {
	*x = RecipeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recipe_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecipeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecipeRequest) ProtoMessage() {}

func (x *RecipeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recipe_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecipeRequest.ProtoReflect.Descriptor instead.
func (*RecipeRequest) Descriptor() ([]byte, []int) {
	return file_recipe_proto_rawDescGZIP(), []int{3}
}

func (x *RecipeRequest) GetRecipe() *Recipe {
	if x != nil {
		return x.Recipe
	}
	return nil
}

type RecipeDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId []byte `protobuf:"bytes,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
}

func (x *RecipeDeleteRequest) Reset() {
	*x = RecipeDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recipe_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecipeDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecipeDeleteRequest) ProtoMessage() {}

func (x *RecipeDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recipe_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecipeDeleteRequest.ProtoReflect.Descriptor instead.
func (*RecipeDeleteRequest) Descriptor() ([]byte, []int) {
	return file_recipe_proto_rawDescGZIP(), []int{4}
}

func (x *RecipeDeleteRequest) GetObjectId() []byte {
	if x != nil {
		return x.ObjectId
	}
	return nil
}

var File_recipe_proto protoreflect.FileDescriptor

var file_recipe_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x6d, 0x65, 0x61, 0x6c, 0x7a, 0x22, 0xb5, 0x01, 0x0a, 0x0a, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64,
	0x69, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x3c, 0x0a, 0x0f, 0x75, 0x6e, 0x69, 0x74, 0x5f,
	0x6f, 0x66, 0x5f, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x14, 0x2e, 0x6d, 0x65, 0x61, 0x6c, 0x7a, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x4f, 0x66, 0x4d,
	0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x52, 0x0d, 0x75, 0x6e, 0x69, 0x74, 0x4f, 0x66, 0x4d, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x76, 0x65, 0x67, 0x65, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x76, 0x65, 0x67, 0x65, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x22, 0xff, 0x01,
	0x0a, 0x06, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x76,
	0x65, 0x67, 0x65, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x76, 0x65, 0x67, 0x65, 0x74, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x12, 0x2e, 0x0a, 0x09, 0x65,
	0x74, 0x68, 0x6e, 0x69, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10,
	0x2e, 0x6d, 0x65, 0x61, 0x6c, 0x7a, 0x2e, 0x45, 0x74, 0x68, 0x6e, 0x69, 0x63, 0x69, 0x74, 0x79,
	0x52, 0x09, 0x65, 0x74, 0x68, 0x6e, 0x69, 0x63, 0x69, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x06, 0x73,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x6d, 0x65,
	0x61, 0x6c, 0x7a, 0x2e, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x73, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x0b, 0x69, 0x6e,
	0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x6d, 0x65, 0x61, 0x6c, 0x7a, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0x2f, 0x0a, 0x10, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x22, 0x36, 0x0a, 0x0d, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x65, 0x61, 0x6c, 0x7a, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65,
	0x52, 0x06, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x22, 0x32, 0x0a, 0x13, 0x52, 0x65, 0x63, 0x69,
	0x70, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x2a, 0x95, 0x01, 0x0a,
	0x09, 0x45, 0x74, 0x68, 0x6e, 0x69, 0x63, 0x69, 0x74, 0x79, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x54, 0x48, 0x4e, 0x49, 0x43, 0x49, 0x54, 0x59, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x4d, 0x45, 0x52, 0x49, 0x43, 0x41, 0x4e, 0x10, 0x01, 0x12,
	0x0d, 0x0a, 0x09, 0x43, 0x41, 0x52, 0x49, 0x42, 0x42, 0x45, 0x41, 0x4e, 0x10, 0x02, 0x12, 0x0b,
	0x0a, 0x07, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x53, 0x45, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x41,
	0x53, 0x49, 0x41, 0x4e, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x54, 0x41, 0x4c, 0x49, 0x41,
	0x4e, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x48, 0x41, 0x49, 0x10, 0x06, 0x12, 0x09, 0x0a,
	0x05, 0x43, 0x41, 0x4a, 0x55, 0x4e, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x4a, 0x41, 0x50, 0x41,
	0x4e, 0x45, 0x53, 0x45, 0x10, 0x08, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x4f, 0x52, 0x4f, 0x43, 0x43,
	0x41, 0x4e, 0x10, 0x09, 0x2a, 0x4a, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x45, 0x41, 0x53, 0x4f, 0x4e,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x57, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x53, 0x50, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x55,
	0x4d, 0x4d, 0x45, 0x52, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x4c, 0x4c, 0x10, 0x04,
	0x2a, 0x72, 0x0a, 0x0d, 0x55, 0x6e, 0x69, 0x74, 0x4f, 0x66, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72,
	0x65, 0x12, 0x19, 0x0a, 0x15, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x55, 0x4e, 0x49,
	0x54, 0x4f, 0x46, 0x4d, 0x45, 0x41, 0x53, 0x55, 0x52, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x54, 0x45, 0x41, 0x53, 0x50, 0x4f, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x41,
	0x42, 0x4c, 0x45, 0x53, 0x50, 0x4f, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x4f,
	0x55, 0x4e, 0x44, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x55, 0x4e, 0x43, 0x45, 0x10, 0x04,
	0x12, 0x07, 0x0a, 0x03, 0x43, 0x55, 0x50, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x49, 0x45,
	0x43, 0x45, 0x10, 0x06, 0x32, 0xd1, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17, 0x2e,
	0x6d, 0x65, 0x61, 0x6c, 0x7a, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x6d, 0x65, 0x61, 0x6c, 0x7a, 0x2e, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12,
	0x14, 0x2e, 0x6d, 0x65, 0x61, 0x6c, 0x7a, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x6d, 0x65, 0x61, 0x6c, 0x7a, 0x2e, 0x52, 0x65,
	0x63, 0x69, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14,
	0x2e, 0x6d, 0x65, 0x61, 0x6c, 0x7a, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x6d, 0x65, 0x61, 0x6c, 0x7a, 0x2e, 0x52, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x2e,
	0x6d, 0x65, 0x61, 0x6c, 0x7a, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x6d, 0x65, 0x61, 0x6c,
	0x7a, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_recipe_proto_rawDescOnce sync.Once
	file_recipe_proto_rawDescData = file_recipe_proto_rawDesc
)

func file_recipe_proto_rawDescGZIP() []byte {
	file_recipe_proto_rawDescOnce.Do(func() {
		file_recipe_proto_rawDescData = protoimpl.X.CompressGZIP(file_recipe_proto_rawDescData)
	})
	return file_recipe_proto_rawDescData
}

var file_recipe_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_recipe_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_recipe_proto_goTypes = []interface{}{
	(Ethnicity)(0),              // 0: mealz.Ethnicity
	(Season)(0),                 // 1: mealz.Season
	(UnitOfMeasure)(0),          // 2: mealz.UnitOfMeasure
	(*Ingredient)(nil),          // 3: mealz.Ingredient
	(*Recipe)(nil),              // 4: mealz.Recipe
	(*RecipeGetRequest)(nil),    // 5: mealz.RecipeGetRequest
	(*RecipeRequest)(nil),       // 6: mealz.RecipeRequest
	(*RecipeDeleteRequest)(nil), // 7: mealz.RecipeDeleteRequest
}
var file_recipe_proto_depIdxs = []int32{
	2, // 0: mealz.Ingredient.unit_of_measure:type_name -> mealz.UnitOfMeasure
	0, // 1: mealz.Recipe.ethnicity:type_name -> mealz.Ethnicity
	1, // 2: mealz.Recipe.season:type_name -> mealz.Season
	3, // 3: mealz.Recipe.ingredients:type_name -> mealz.Ingredient
	4, // 4: mealz.RecipeRequest.recipe:type_name -> mealz.Recipe
	5, // 5: mealz.RecipeService.Get:input_type -> mealz.RecipeGetRequest
	6, // 6: mealz.RecipeService.Insert:input_type -> mealz.RecipeRequest
	6, // 7: mealz.RecipeService.Update:input_type -> mealz.RecipeRequest
	7, // 8: mealz.RecipeService.Delete:input_type -> mealz.RecipeDeleteRequest
	4, // 9: mealz.RecipeService.Get:output_type -> mealz.Recipe
	4, // 10: mealz.RecipeService.Insert:output_type -> mealz.Recipe
	4, // 11: mealz.RecipeService.Update:output_type -> mealz.Recipe
	4, // 12: mealz.RecipeService.Delete:output_type -> mealz.Recipe
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_recipe_proto_init() }
func file_recipe_proto_init() {
	if File_recipe_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_recipe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ingredient); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_recipe_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Recipe); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_recipe_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecipeGetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_recipe_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecipeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_recipe_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecipeDeleteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_recipe_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_recipe_proto_goTypes,
		DependencyIndexes: file_recipe_proto_depIdxs,
		EnumInfos:         file_recipe_proto_enumTypes,
		MessageInfos:      file_recipe_proto_msgTypes,
	}.Build()
	File_recipe_proto = out.File
	file_recipe_proto_rawDesc = nil
	file_recipe_proto_goTypes = nil
	file_recipe_proto_depIdxs = nil
}
