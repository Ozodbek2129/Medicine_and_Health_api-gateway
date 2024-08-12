// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.18.0
// source: Medicine_and_Health_protos/HealthAnalytics/Health_Analytics.proto

package health_analytics

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	HealthAnalyticsService_AddMedicalRecord_FullMethodName              = "/healthanalytics.HealthAnalyticsService/AddMedicalRecord"
	HealthAnalyticsService_GetMedicalRecord_FullMethodName              = "/healthanalytics.HealthAnalyticsService/GetMedicalRecord"
	HealthAnalyticsService_UpdateMedicalRecord_FullMethodName           = "/healthanalytics.HealthAnalyticsService/UpdateMedicalRecord"
	HealthAnalyticsService_DeleteMedicalRecord_FullMethodName           = "/healthanalytics.HealthAnalyticsService/DeleteMedicalRecord"
	HealthAnalyticsService_ListMedicalRecords_FullMethodName            = "/healthanalytics.HealthAnalyticsService/ListMedicalRecords"
	HealthAnalyticsService_AddLifestyleData_FullMethodName              = "/healthanalytics.HealthAnalyticsService/AddLifestyleData"
	HealthAnalyticsService_GetLifestyleData_FullMethodName              = "/healthanalytics.HealthAnalyticsService/GetLifestyleData"
	HealthAnalyticsService_GetAllLifestyleData_FullMethodName           = "/healthanalytics.HealthAnalyticsService/GetAllLifestyleData"
	HealthAnalyticsService_UpdateLifestyleData_FullMethodName           = "/healthanalytics.HealthAnalyticsService/UpdateLifestyleData"
	HealthAnalyticsService_DeleteLifestyleData_FullMethodName           = "/healthanalytics.HealthAnalyticsService/DeleteLifestyleData"
	HealthAnalyticsService_AddWearableData_FullMethodName               = "/healthanalytics.HealthAnalyticsService/AddWearableData"
	HealthAnalyticsService_GetWearableData_FullMethodName               = "/healthanalytics.HealthAnalyticsService/GetWearableData"
	HealthAnalyticsService_GetAllWearableData_FullMethodName            = "/healthanalytics.HealthAnalyticsService/GetAllWearableData"
	HealthAnalyticsService_UpdateWearableData_FullMethodName            = "/healthanalytics.HealthAnalyticsService/UpdateWearableData"
	HealthAnalyticsService_DeleteWearableData_FullMethodName            = "/healthanalytics.HealthAnalyticsService/DeleteWearableData"
	HealthAnalyticsService_GenerateHealthRecommendations_FullMethodName = "/healthanalytics.HealthAnalyticsService/GenerateHealthRecommendations"
	HealthAnalyticsService_GetRealtimeHealthMonitoring_FullMethodName   = "/healthanalytics.HealthAnalyticsService/GetRealtimeHealthMonitoring"
	HealthAnalyticsService_GetDailyHealthSummary_FullMethodName         = "/healthanalytics.HealthAnalyticsService/GetDailyHealthSummary"
	HealthAnalyticsService_GetWeeklyHealthSummary_FullMethodName        = "/healthanalytics.HealthAnalyticsService/GetWeeklyHealthSummary"
)

// HealthAnalyticsServiceClient is the client API for HealthAnalyticsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Health Analytics Service uchun servis ta'rifi
type HealthAnalyticsServiceClient interface {
	// Tibbiy yozuvlar uchun RPC lar
	AddMedicalRecord(ctx context.Context, in *AddMedicalRecordRequest, opts ...grpc.CallOption) (*AddMedicalRecordResponse, error)
	GetMedicalRecord(ctx context.Context, in *GetMedicalRecordRequest, opts ...grpc.CallOption) (*GetMedicalRecordResponse, error)
	UpdateMedicalRecord(ctx context.Context, in *UpdateMedicalRecordRequest, opts ...grpc.CallOption) (*UpdateMedicalRecordResponse, error)
	DeleteMedicalRecord(ctx context.Context, in *DeleteMedicalRecordRequest, opts ...grpc.CallOption) (*DeleteMedicalRecordResponse, error)
	ListMedicalRecords(ctx context.Context, in *ListMedicalRecordsRequest, opts ...grpc.CallOption) (*ListMedicalRecordsResponse, error)
	// Turmush tarzi ma'lumotlari uchun RPC lar
	AddLifestyleData(ctx context.Context, in *AddLifestyleDataRequest, opts ...grpc.CallOption) (*AddLifestyleDataResponse, error)
	GetLifestyleData(ctx context.Context, in *GetLifestyleDataRequest, opts ...grpc.CallOption) (*GetLifestyleDataResponse, error)
	GetAllLifestyleData(ctx context.Context, in *GetAllLifestyleDataRequest, opts ...grpc.CallOption) (*GetAllLifestyleDataResponse, error)
	UpdateLifestyleData(ctx context.Context, in *UpdateLifestyleDataRequest, opts ...grpc.CallOption) (*UpdateLifestyleDataResponse, error)
	DeleteLifestyleData(ctx context.Context, in *DeleteLifestyleDataRequest, opts ...grpc.CallOption) (*DeleteLifestyleDataResponse, error)
	// Kiyiladigan qurilma ma'lumotlari uchun RPC lar
	AddWearableData(ctx context.Context, in *AddWearableDataRequest, opts ...grpc.CallOption) (*AddWearableDataResponse, error)
	GetWearableData(ctx context.Context, in *GetWearableDataRequest, opts ...grpc.CallOption) (*GetWearableDataResponse, error)
	GetAllWearableData(ctx context.Context, in *GetAllWearableDataRequest, opts ...grpc.CallOption) (*GetAllWearableDataResponse, error)
	UpdateWearableData(ctx context.Context, in *UpdateWearableDataRequest, opts ...grpc.CallOption) (*UpdateWearableDataResponse, error)
	DeleteWearableData(ctx context.Context, in *DeleteWearableDataRequest, opts ...grpc.CallOption) (*DeleteWearableDataResponse, error)
	// Sog'liq tavsiyalari va monitoringi uchun RPC lar
	GenerateHealthRecommendations(ctx context.Context, in *GenerateHealthRecommendationsRequest, opts ...grpc.CallOption) (*GenerateHealthRecommendationsResponse, error)
	GetRealtimeHealthMonitoring(ctx context.Context, in *GetRealtimeHealthMonitoringRequest, opts ...grpc.CallOption) (*GetRealtimeHealthMonitoringResponse, error)
	GetDailyHealthSummary(ctx context.Context, in *GetDailyHealthSummaryRequest, opts ...grpc.CallOption) (*GetDailyHealthSummaryResponse, error)
	GetWeeklyHealthSummary(ctx context.Context, in *GetWeeklyHealthSummaryRequest, opts ...grpc.CallOption) (*GetWeeklyHealthSummaryResponse, error)
}

type healthAnalyticsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthAnalyticsServiceClient(cc grpc.ClientConnInterface) HealthAnalyticsServiceClient {
	return &healthAnalyticsServiceClient{cc}
}

func (c *healthAnalyticsServiceClient) AddMedicalRecord(ctx context.Context, in *AddMedicalRecordRequest, opts ...grpc.CallOption) (*AddMedicalRecordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddMedicalRecordResponse)
	err := c.cc.Invoke(ctx, HealthAnalyticsService_AddMedicalRecord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) GetMedicalRecord(ctx context.Context, in *GetMedicalRecordRequest, opts ...grpc.CallOption) (*GetMedicalRecordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMedicalRecordResponse)
	err := c.cc.Invoke(ctx, HealthAnalyticsService_GetMedicalRecord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) UpdateMedicalRecord(ctx context.Context, in *UpdateMedicalRecordRequest, opts ...grpc.CallOption) (*UpdateMedicalRecordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateMedicalRecordResponse)
	err := c.cc.Invoke(ctx, HealthAnalyticsService_UpdateMedicalRecord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) DeleteMedicalRecord(ctx context.Context, in *DeleteMedicalRecordRequest, opts ...grpc.CallOption) (*DeleteMedicalRecordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteMedicalRecordResponse)
	err := c.cc.Invoke(ctx, HealthAnalyticsService_DeleteMedicalRecord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) ListMedicalRecords(ctx context.Context, in *ListMedicalRecordsRequest, opts ...grpc.CallOption) (*ListMedicalRecordsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListMedicalRecordsResponse)
	err := c.cc.Invoke(ctx, HealthAnalyticsService_ListMedicalRecords_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) AddLifestyleData(ctx context.Context, in *AddLifestyleDataRequest, opts ...grpc.CallOption) (*AddLifestyleDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddLifestyleDataResponse)
	err := c.cc.Invoke(ctx, HealthAnalyticsService_AddLifestyleData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) GetLifestyleData(ctx context.Context, in *GetLifestyleDataRequest, opts ...grpc.CallOption) (*GetLifestyleDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLifestyleDataResponse)
	err := c.cc.Invoke(ctx, HealthAnalyticsService_GetLifestyleData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) GetAllLifestyleData(ctx context.Context, in *GetAllLifestyleDataRequest, opts ...grpc.CallOption) (*GetAllLifestyleDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllLifestyleDataResponse)
	err := c.cc.Invoke(ctx, HealthAnalyticsService_GetAllLifestyleData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) UpdateLifestyleData(ctx context.Context, in *UpdateLifestyleDataRequest, opts ...grpc.CallOption) (*UpdateLifestyleDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateLifestyleDataResponse)
	err := c.cc.Invoke(ctx, HealthAnalyticsService_UpdateLifestyleData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) DeleteLifestyleData(ctx context.Context, in *DeleteLifestyleDataRequest, opts ...grpc.CallOption) (*DeleteLifestyleDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteLifestyleDataResponse)
	err := c.cc.Invoke(ctx, HealthAnalyticsService_DeleteLifestyleData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) AddWearableData(ctx context.Context, in *AddWearableDataRequest, opts ...grpc.CallOption) (*AddWearableDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddWearableDataResponse)
	err := c.cc.Invoke(ctx, HealthAnalyticsService_AddWearableData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) GetWearableData(ctx context.Context, in *GetWearableDataRequest, opts ...grpc.CallOption) (*GetWearableDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWearableDataResponse)
	err := c.cc.Invoke(ctx, HealthAnalyticsService_GetWearableData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) GetAllWearableData(ctx context.Context, in *GetAllWearableDataRequest, opts ...grpc.CallOption) (*GetAllWearableDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllWearableDataResponse)
	err := c.cc.Invoke(ctx, HealthAnalyticsService_GetAllWearableData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) UpdateWearableData(ctx context.Context, in *UpdateWearableDataRequest, opts ...grpc.CallOption) (*UpdateWearableDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateWearableDataResponse)
	err := c.cc.Invoke(ctx, HealthAnalyticsService_UpdateWearableData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) DeleteWearableData(ctx context.Context, in *DeleteWearableDataRequest, opts ...grpc.CallOption) (*DeleteWearableDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteWearableDataResponse)
	err := c.cc.Invoke(ctx, HealthAnalyticsService_DeleteWearableData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) GenerateHealthRecommendations(ctx context.Context, in *GenerateHealthRecommendationsRequest, opts ...grpc.CallOption) (*GenerateHealthRecommendationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateHealthRecommendationsResponse)
	err := c.cc.Invoke(ctx, HealthAnalyticsService_GenerateHealthRecommendations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) GetRealtimeHealthMonitoring(ctx context.Context, in *GetRealtimeHealthMonitoringRequest, opts ...grpc.CallOption) (*GetRealtimeHealthMonitoringResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRealtimeHealthMonitoringResponse)
	err := c.cc.Invoke(ctx, HealthAnalyticsService_GetRealtimeHealthMonitoring_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) GetDailyHealthSummary(ctx context.Context, in *GetDailyHealthSummaryRequest, opts ...grpc.CallOption) (*GetDailyHealthSummaryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDailyHealthSummaryResponse)
	err := c.cc.Invoke(ctx, HealthAnalyticsService_GetDailyHealthSummary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthAnalyticsServiceClient) GetWeeklyHealthSummary(ctx context.Context, in *GetWeeklyHealthSummaryRequest, opts ...grpc.CallOption) (*GetWeeklyHealthSummaryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWeeklyHealthSummaryResponse)
	err := c.cc.Invoke(ctx, HealthAnalyticsService_GetWeeklyHealthSummary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthAnalyticsServiceServer is the server API for HealthAnalyticsService service.
// All implementations must embed UnimplementedHealthAnalyticsServiceServer
// for forward compatibility
//
// Health Analytics Service uchun servis ta'rifi
type HealthAnalyticsServiceServer interface {
	// Tibbiy yozuvlar uchun RPC lar
	AddMedicalRecord(context.Context, *AddMedicalRecordRequest) (*AddMedicalRecordResponse, error)
	GetMedicalRecord(context.Context, *GetMedicalRecordRequest) (*GetMedicalRecordResponse, error)
	UpdateMedicalRecord(context.Context, *UpdateMedicalRecordRequest) (*UpdateMedicalRecordResponse, error)
	DeleteMedicalRecord(context.Context, *DeleteMedicalRecordRequest) (*DeleteMedicalRecordResponse, error)
	ListMedicalRecords(context.Context, *ListMedicalRecordsRequest) (*ListMedicalRecordsResponse, error)
	// Turmush tarzi ma'lumotlari uchun RPC lar
	AddLifestyleData(context.Context, *AddLifestyleDataRequest) (*AddLifestyleDataResponse, error)
	GetLifestyleData(context.Context, *GetLifestyleDataRequest) (*GetLifestyleDataResponse, error)
	GetAllLifestyleData(context.Context, *GetAllLifestyleDataRequest) (*GetAllLifestyleDataResponse, error)
	UpdateLifestyleData(context.Context, *UpdateLifestyleDataRequest) (*UpdateLifestyleDataResponse, error)
	DeleteLifestyleData(context.Context, *DeleteLifestyleDataRequest) (*DeleteLifestyleDataResponse, error)
	// Kiyiladigan qurilma ma'lumotlari uchun RPC lar
	AddWearableData(context.Context, *AddWearableDataRequest) (*AddWearableDataResponse, error)
	GetWearableData(context.Context, *GetWearableDataRequest) (*GetWearableDataResponse, error)
	GetAllWearableData(context.Context, *GetAllWearableDataRequest) (*GetAllWearableDataResponse, error)
	UpdateWearableData(context.Context, *UpdateWearableDataRequest) (*UpdateWearableDataResponse, error)
	DeleteWearableData(context.Context, *DeleteWearableDataRequest) (*DeleteWearableDataResponse, error)
	// Sog'liq tavsiyalari va monitoringi uchun RPC lar
	GenerateHealthRecommendations(context.Context, *GenerateHealthRecommendationsRequest) (*GenerateHealthRecommendationsResponse, error)
	GetRealtimeHealthMonitoring(context.Context, *GetRealtimeHealthMonitoringRequest) (*GetRealtimeHealthMonitoringResponse, error)
	GetDailyHealthSummary(context.Context, *GetDailyHealthSummaryRequest) (*GetDailyHealthSummaryResponse, error)
	GetWeeklyHealthSummary(context.Context, *GetWeeklyHealthSummaryRequest) (*GetWeeklyHealthSummaryResponse, error)
	mustEmbedUnimplementedHealthAnalyticsServiceServer()
}

// UnimplementedHealthAnalyticsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHealthAnalyticsServiceServer struct {
}

func (UnimplementedHealthAnalyticsServiceServer) AddMedicalRecord(context.Context, *AddMedicalRecordRequest) (*AddMedicalRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMedicalRecord not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) GetMedicalRecord(context.Context, *GetMedicalRecordRequest) (*GetMedicalRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMedicalRecord not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) UpdateMedicalRecord(context.Context, *UpdateMedicalRecordRequest) (*UpdateMedicalRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMedicalRecord not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) DeleteMedicalRecord(context.Context, *DeleteMedicalRecordRequest) (*DeleteMedicalRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMedicalRecord not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) ListMedicalRecords(context.Context, *ListMedicalRecordsRequest) (*ListMedicalRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMedicalRecords not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) AddLifestyleData(context.Context, *AddLifestyleDataRequest) (*AddLifestyleDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLifestyleData not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) GetLifestyleData(context.Context, *GetLifestyleDataRequest) (*GetLifestyleDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLifestyleData not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) GetAllLifestyleData(context.Context, *GetAllLifestyleDataRequest) (*GetAllLifestyleDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllLifestyleData not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) UpdateLifestyleData(context.Context, *UpdateLifestyleDataRequest) (*UpdateLifestyleDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLifestyleData not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) DeleteLifestyleData(context.Context, *DeleteLifestyleDataRequest) (*DeleteLifestyleDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLifestyleData not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) AddWearableData(context.Context, *AddWearableDataRequest) (*AddWearableDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddWearableData not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) GetWearableData(context.Context, *GetWearableDataRequest) (*GetWearableDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWearableData not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) GetAllWearableData(context.Context, *GetAllWearableDataRequest) (*GetAllWearableDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllWearableData not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) UpdateWearableData(context.Context, *UpdateWearableDataRequest) (*UpdateWearableDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWearableData not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) DeleteWearableData(context.Context, *DeleteWearableDataRequest) (*DeleteWearableDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWearableData not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) GenerateHealthRecommendations(context.Context, *GenerateHealthRecommendationsRequest) (*GenerateHealthRecommendationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateHealthRecommendations not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) GetRealtimeHealthMonitoring(context.Context, *GetRealtimeHealthMonitoringRequest) (*GetRealtimeHealthMonitoringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRealtimeHealthMonitoring not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) GetDailyHealthSummary(context.Context, *GetDailyHealthSummaryRequest) (*GetDailyHealthSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDailyHealthSummary not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) GetWeeklyHealthSummary(context.Context, *GetWeeklyHealthSummaryRequest) (*GetWeeklyHealthSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeeklyHealthSummary not implemented")
}
func (UnimplementedHealthAnalyticsServiceServer) mustEmbedUnimplementedHealthAnalyticsServiceServer() {
}

// UnsafeHealthAnalyticsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthAnalyticsServiceServer will
// result in compilation errors.
type UnsafeHealthAnalyticsServiceServer interface {
	mustEmbedUnimplementedHealthAnalyticsServiceServer()
}

func RegisterHealthAnalyticsServiceServer(s grpc.ServiceRegistrar, srv HealthAnalyticsServiceServer) {
	s.RegisterService(&HealthAnalyticsService_ServiceDesc, srv)
}

func _HealthAnalyticsService_AddMedicalRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMedicalRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).AddMedicalRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthAnalyticsService_AddMedicalRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).AddMedicalRecord(ctx, req.(*AddMedicalRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_GetMedicalRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMedicalRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).GetMedicalRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthAnalyticsService_GetMedicalRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).GetMedicalRecord(ctx, req.(*GetMedicalRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_UpdateMedicalRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMedicalRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).UpdateMedicalRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthAnalyticsService_UpdateMedicalRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).UpdateMedicalRecord(ctx, req.(*UpdateMedicalRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_DeleteMedicalRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMedicalRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).DeleteMedicalRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthAnalyticsService_DeleteMedicalRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).DeleteMedicalRecord(ctx, req.(*DeleteMedicalRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_ListMedicalRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMedicalRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).ListMedicalRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthAnalyticsService_ListMedicalRecords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).ListMedicalRecords(ctx, req.(*ListMedicalRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_AddLifestyleData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLifestyleDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).AddLifestyleData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthAnalyticsService_AddLifestyleData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).AddLifestyleData(ctx, req.(*AddLifestyleDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_GetLifestyleData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLifestyleDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).GetLifestyleData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthAnalyticsService_GetLifestyleData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).GetLifestyleData(ctx, req.(*GetLifestyleDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_GetAllLifestyleData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllLifestyleDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).GetAllLifestyleData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthAnalyticsService_GetAllLifestyleData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).GetAllLifestyleData(ctx, req.(*GetAllLifestyleDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_UpdateLifestyleData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLifestyleDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).UpdateLifestyleData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthAnalyticsService_UpdateLifestyleData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).UpdateLifestyleData(ctx, req.(*UpdateLifestyleDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_DeleteLifestyleData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLifestyleDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).DeleteLifestyleData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthAnalyticsService_DeleteLifestyleData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).DeleteLifestyleData(ctx, req.(*DeleteLifestyleDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_AddWearableData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddWearableDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).AddWearableData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthAnalyticsService_AddWearableData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).AddWearableData(ctx, req.(*AddWearableDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_GetWearableData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWearableDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).GetWearableData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthAnalyticsService_GetWearableData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).GetWearableData(ctx, req.(*GetWearableDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_GetAllWearableData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllWearableDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).GetAllWearableData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthAnalyticsService_GetAllWearableData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).GetAllWearableData(ctx, req.(*GetAllWearableDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_UpdateWearableData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWearableDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).UpdateWearableData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthAnalyticsService_UpdateWearableData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).UpdateWearableData(ctx, req.(*UpdateWearableDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_DeleteWearableData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWearableDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).DeleteWearableData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthAnalyticsService_DeleteWearableData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).DeleteWearableData(ctx, req.(*DeleteWearableDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_GenerateHealthRecommendations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateHealthRecommendationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).GenerateHealthRecommendations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthAnalyticsService_GenerateHealthRecommendations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).GenerateHealthRecommendations(ctx, req.(*GenerateHealthRecommendationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_GetRealtimeHealthMonitoring_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRealtimeHealthMonitoringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).GetRealtimeHealthMonitoring(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthAnalyticsService_GetRealtimeHealthMonitoring_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).GetRealtimeHealthMonitoring(ctx, req.(*GetRealtimeHealthMonitoringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_GetDailyHealthSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDailyHealthSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).GetDailyHealthSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthAnalyticsService_GetDailyHealthSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).GetDailyHealthSummary(ctx, req.(*GetDailyHealthSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthAnalyticsService_GetWeeklyHealthSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWeeklyHealthSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthAnalyticsServiceServer).GetWeeklyHealthSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthAnalyticsService_GetWeeklyHealthSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthAnalyticsServiceServer).GetWeeklyHealthSummary(ctx, req.(*GetWeeklyHealthSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HealthAnalyticsService_ServiceDesc is the grpc.ServiceDesc for HealthAnalyticsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HealthAnalyticsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "healthanalytics.HealthAnalyticsService",
	HandlerType: (*HealthAnalyticsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMedicalRecord",
			Handler:    _HealthAnalyticsService_AddMedicalRecord_Handler,
		},
		{
			MethodName: "GetMedicalRecord",
			Handler:    _HealthAnalyticsService_GetMedicalRecord_Handler,
		},
		{
			MethodName: "UpdateMedicalRecord",
			Handler:    _HealthAnalyticsService_UpdateMedicalRecord_Handler,
		},
		{
			MethodName: "DeleteMedicalRecord",
			Handler:    _HealthAnalyticsService_DeleteMedicalRecord_Handler,
		},
		{
			MethodName: "ListMedicalRecords",
			Handler:    _HealthAnalyticsService_ListMedicalRecords_Handler,
		},
		{
			MethodName: "AddLifestyleData",
			Handler:    _HealthAnalyticsService_AddLifestyleData_Handler,
		},
		{
			MethodName: "GetLifestyleData",
			Handler:    _HealthAnalyticsService_GetLifestyleData_Handler,
		},
		{
			MethodName: "GetAllLifestyleData",
			Handler:    _HealthAnalyticsService_GetAllLifestyleData_Handler,
		},
		{
			MethodName: "UpdateLifestyleData",
			Handler:    _HealthAnalyticsService_UpdateLifestyleData_Handler,
		},
		{
			MethodName: "DeleteLifestyleData",
			Handler:    _HealthAnalyticsService_DeleteLifestyleData_Handler,
		},
		{
			MethodName: "AddWearableData",
			Handler:    _HealthAnalyticsService_AddWearableData_Handler,
		},
		{
			MethodName: "GetWearableData",
			Handler:    _HealthAnalyticsService_GetWearableData_Handler,
		},
		{
			MethodName: "GetAllWearableData",
			Handler:    _HealthAnalyticsService_GetAllWearableData_Handler,
		},
		{
			MethodName: "UpdateWearableData",
			Handler:    _HealthAnalyticsService_UpdateWearableData_Handler,
		},
		{
			MethodName: "DeleteWearableData",
			Handler:    _HealthAnalyticsService_DeleteWearableData_Handler,
		},
		{
			MethodName: "GenerateHealthRecommendations",
			Handler:    _HealthAnalyticsService_GenerateHealthRecommendations_Handler,
		},
		{
			MethodName: "GetRealtimeHealthMonitoring",
			Handler:    _HealthAnalyticsService_GetRealtimeHealthMonitoring_Handler,
		},
		{
			MethodName: "GetDailyHealthSummary",
			Handler:    _HealthAnalyticsService_GetDailyHealthSummary_Handler,
		},
		{
			MethodName: "GetWeeklyHealthSummary",
			Handler:    _HealthAnalyticsService_GetWeeklyHealthSummary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Medicine_and_Health_protos/HealthAnalytics/Health_Analytics.proto",
}
