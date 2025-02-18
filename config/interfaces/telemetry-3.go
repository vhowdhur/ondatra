package interfaces

// This file contains generated telemetry method augmentations for the
// generated path structs, which makes use of their gNMI paths for making
// ONDATRA telemetry calls.

import (
	"reflect"
	"testing"

	config "github.com/openconfig/ondatra/config"
	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	oc "github.com/openconfig/ondatra/telemetry"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/config/subinterface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/config/subinterface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/config/subinterface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/config/subinterface with a ONCE subscription.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/config/subinterface.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/config/subinterface in the given batch object.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/config/subinterface.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/config/subinterface in the given batch object.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/config/subinterface.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv4/unnumbered/interface-ref/config/subinterface in the given batch object.
func (n *Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath extracts the value of the leaf Subinterface from its parent oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_Subinterface_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv4_Unnumbered_InterfaceRef) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Subinterface
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6Path) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Ipv6 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Ipv6{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6 with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6Path) Get(t testing.TB) *oc.Interface_Subinterface_Ipv6 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6PathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Ipv6 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Ipv6
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Ipv6{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6 with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6PathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Ipv6 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Ipv6
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6.
func (n *Interface_Subinterface_Ipv6Path) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6 in the given batch object.
func (n *Interface_Subinterface_Ipv6Path) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6.
func (n *Interface_Subinterface_Ipv6Path) Replace(t testing.TB, val *oc.Interface_Subinterface_Ipv6) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6 in the given batch object.
func (n *Interface_Subinterface_Ipv6Path) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Ipv6) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6.
func (n *Interface_Subinterface_Ipv6Path) Update(t testing.TB, val *oc.Interface_Subinterface_Ipv6) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6 in the given batch object.
func (n *Interface_Subinterface_Ipv6Path) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Ipv6) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_AddressPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Ipv6_Address {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Ipv6_Address{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_AddressPath) Get(t testing.TB) *oc.Interface_Subinterface_Ipv6_Address {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Ipv6_Address {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Ipv6_Address
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Ipv6_Address{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_AddressPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Ipv6_Address {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Ipv6_Address
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address.
func (n *Interface_Subinterface_Ipv6_AddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address in the given batch object.
func (n *Interface_Subinterface_Ipv6_AddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address.
func (n *Interface_Subinterface_Ipv6_AddressPath) Replace(t testing.TB, val *oc.Interface_Subinterface_Ipv6_Address) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address in the given batch object.
func (n *Interface_Subinterface_Ipv6_AddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Ipv6_Address) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address.
func (n *Interface_Subinterface_Ipv6_AddressPath) Update(t testing.TB, val *oc.Interface_Subinterface_Ipv6_Address) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address in the given batch object.
func (n *Interface_Subinterface_Ipv6_AddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Ipv6_Address) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/config/ip with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_IpPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_IpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/config/ip with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_IpPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/config/ip with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_IpPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_IpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/config/ip with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_IpPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/config/ip.
func (n *Interface_Subinterface_Ipv6_Address_IpPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/config/ip in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_IpPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/config/ip.
func (n *Interface_Subinterface_Ipv6_Address_IpPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/config/ip in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_IpPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/config/ip.
func (n *Interface_Subinterface_Ipv6_Address_IpPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/config/ip in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_IpPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv6_Address_IpPath extracts the value of the leaf Ip from its parent oc.Interface_Subinterface_Ipv6_Address
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_Subinterface_Ipv6_Address_IpPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Ip
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/config/prefix-length with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_PrefixLengthPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_PrefixLengthPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/config/prefix-length with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_PrefixLengthPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/config/prefix-length with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_PrefixLengthPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_PrefixLengthPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/config/prefix-length with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_PrefixLengthPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/config/prefix-length.
func (n *Interface_Subinterface_Ipv6_Address_PrefixLengthPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/config/prefix-length in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_PrefixLengthPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/config/prefix-length.
func (n *Interface_Subinterface_Ipv6_Address_PrefixLengthPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/config/prefix-length in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_PrefixLengthPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/config/prefix-length.
func (n *Interface_Subinterface_Ipv6_Address_PrefixLengthPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/config/prefix-length in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_PrefixLengthPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv6_Address_PrefixLengthPath extracts the value of the leaf PrefixLength from its parent oc.Interface_Subinterface_Ipv6_Address
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertInterface_Subinterface_Ipv6_Address_PrefixLengthPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.PrefixLength
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroupPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroupPath) Get(t testing.TB) *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroupPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroupPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Ipv6_Address_VrrpGroup {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Ipv6_Address_VrrpGroup
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroupPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroupPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroupPath) Replace(t testing.TB, val *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroupPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroupPath) Update(t testing.TB, val *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroupPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/accept-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetAcceptMode())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/accept-mode with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/accept-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/accept-mode with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/accept-mode.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/accept-mode in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/accept-mode.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/accept-mode in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/accept-mode.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/accept-mode in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePath extracts the value of the leaf AcceptMode from its parent oc.Interface_Subinterface_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_Subinterface_Ipv6_Address_VrrpGroup_AcceptModePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.AcceptMode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/advertisement-interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint16{
		Metadata: md,
	}).SetVal(goStruct.GetAdvertisementInterval())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/advertisement-interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/advertisement-interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/advertisement-interval with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/advertisement-interval.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/advertisement-interval in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/advertisement-interval.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/advertisement-interval in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/advertisement-interval.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/advertisement-interval in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath extracts the value of the leaf AdvertisementInterval from its parent oc.Interface_Subinterface_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.AdvertisementInterval
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) Get(t testing.TB) *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) Replace(t testing.TB, val *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) Update(t testing.TB, val *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint8{
		Metadata: md,
	}).SetVal(goStruct.GetPriorityDecrement())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/priority-decrement in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath extracts the value of the leaf PriorityDecrement from its parent oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.PriorityDecrement
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Replace(t testing.TB, val []string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []string) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) Update(t testing.TB, val []string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/config/track-interface in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []string) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath extracts the value of the leaf TrackInterface from its parent oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertInterface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking_TrackInterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup_InterfaceTracking) *oc.QualifiedStringSlice {
	t.Helper()
	qv := &oc.QualifiedStringSlice{
		Metadata: md,
	}
	val := parent.TrackInterface
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/preempt-delay with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint16{
		Metadata: md,
	}).SetVal(goStruct.GetPreemptDelay())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/preempt-delay with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/preempt-delay with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/preempt-delay with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/preempt-delay.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/preempt-delay in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/preempt-delay.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/preempt-delay in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/preempt-delay.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/preempt-delay in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPath extracts the value of the leaf PreemptDelay from its parent oc.Interface_Subinterface_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PreemptDelayPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.PreemptDelay
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/preempt with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetPreempt())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/preempt with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/preempt with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/preempt with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/preempt.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/preempt in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/preempt.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/preempt in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/preempt.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/preempt in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPath extracts the value of the leaf Preempt from its parent oc.Interface_Subinterface_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PreemptPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Preempt
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/priority with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint8{
		Metadata: md,
	}).SetVal(goStruct.GetPriority())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/priority with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/priority with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/priority with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/priority.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/priority in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/priority.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/priority in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/priority.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/priority in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPath extracts the value of the leaf Priority from its parent oc.Interface_Subinterface_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertInterface_Subinterface_Ipv6_Address_VrrpGroup_PriorityPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.Priority
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPath) Lookup(t testing.TB) *oc.QualifiedStringSlice {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPath) Get(t testing.TB) []string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedStringSlice {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedStringSlice
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-address with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPathAny) Get(t testing.TB) [][]string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data [][]string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-address.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-address in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-address.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPath) Replace(t testing.TB, val []string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-address in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val []string) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-address.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPath) Update(t testing.TB, val []string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-address in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val []string) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPath extracts the value of the leaf VirtualAddress from its parent oc.Interface_Subinterface_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedStringSlice.
func convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup) *oc.QualifiedStringSlice {
	t.Helper()
	qv := &oc.QualifiedStringSlice{
		Metadata: md,
	}
	val := parent.VirtualAddress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-link-local with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-link-local with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-link-local with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-link-local with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-link-local.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-link-local in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-link-local.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-link-local in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-link-local.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-link-local in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath extracts the value of the leaf VirtualLinkLocal from its parent oc.Interface_Subinterface_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualLinkLocalPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.VirtualLinkLocal
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-router-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-router-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-router-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-router-id with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-router-id.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-router-id in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-router-id.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) Replace(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-router-id in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-router-id.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) Update(t testing.TB, val uint8) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/addresses/address/vrrp/vrrp-group/config/virtual-router-id in the given batch object.
func (n *Interface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint8) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPath extracts the value of the leaf VirtualRouterId from its parent oc.Interface_Subinterface_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertInterface_Subinterface_Ipv6_Address_VrrpGroup_VirtualRouterIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Address_VrrpGroup) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.VirtualRouterId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_AutoconfPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Ipv6_Autoconf {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Autoconf{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Autoconf", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Ipv6_Autoconf{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_AutoconfPath) Get(t testing.TB) *oc.Interface_Subinterface_Ipv6_Autoconf {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_AutoconfPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Ipv6_Autoconf {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Ipv6_Autoconf
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Autoconf{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Autoconf", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Ipv6_Autoconf{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_AutoconfPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Ipv6_Autoconf {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Ipv6_Autoconf
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf.
func (n *Interface_Subinterface_Ipv6_AutoconfPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf in the given batch object.
func (n *Interface_Subinterface_Ipv6_AutoconfPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf.
func (n *Interface_Subinterface_Ipv6_AutoconfPath) Replace(t testing.TB, val *oc.Interface_Subinterface_Ipv6_Autoconf) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf in the given batch object.
func (n *Interface_Subinterface_Ipv6_AutoconfPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Ipv6_Autoconf) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf.
func (n *Interface_Subinterface_Ipv6_AutoconfPath) Update(t testing.TB, val *oc.Interface_Subinterface_Ipv6_Autoconf) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf in the given batch object.
func (n *Interface_Subinterface_Ipv6_AutoconfPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Ipv6_Autoconf) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/create-global-addresses with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Autoconf{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Autoconf", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetCreateGlobalAddresses())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/create-global-addresses with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/create-global-addresses with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Autoconf{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Autoconf", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/create-global-addresses with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/create-global-addresses.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/create-global-addresses in the given batch object.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/create-global-addresses.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/create-global-addresses in the given batch object.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/create-global-addresses.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/create-global-addresses in the given batch object.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPath extracts the value of the leaf CreateGlobalAddresses from its parent oc.Interface_Subinterface_Ipv6_Autoconf
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_Subinterface_Ipv6_Autoconf_CreateGlobalAddressesPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Autoconf) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.CreateGlobalAddresses
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/create-temporary-addresses with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Autoconf{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Autoconf", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetCreateTemporaryAddresses())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/create-temporary-addresses with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/create-temporary-addresses with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Autoconf{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Autoconf", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/create-temporary-addresses with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/create-temporary-addresses.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/create-temporary-addresses in the given batch object.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/create-temporary-addresses.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/create-temporary-addresses in the given batch object.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/create-temporary-addresses.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/create-temporary-addresses in the given batch object.
func (n *Interface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPath extracts the value of the leaf CreateTemporaryAddresses from its parent oc.Interface_Subinterface_Ipv6_Autoconf
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_Subinterface_Ipv6_Autoconf_CreateTemporaryAddressesPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Autoconf) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.CreateTemporaryAddresses
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/temporary-preferred-lifetime with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Autoconf{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Autoconf", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePath(t, md, goStruct)
	}
	return (&oc.QualifiedUint32{
		Metadata: md,
	}).SetVal(goStruct.GetTemporaryPreferredLifetime())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/temporary-preferred-lifetime with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/temporary-preferred-lifetime with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Autoconf{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Autoconf", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/temporary-preferred-lifetime with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/temporary-preferred-lifetime.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/temporary-preferred-lifetime in the given batch object.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/temporary-preferred-lifetime.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/temporary-preferred-lifetime in the given batch object.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/temporary-preferred-lifetime.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/temporary-preferred-lifetime in the given batch object.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePath extracts the value of the leaf TemporaryPreferredLifetime from its parent oc.Interface_Subinterface_Ipv6_Autoconf
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_Subinterface_Ipv6_Autoconf_TemporaryPreferredLifetimePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Autoconf) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.TemporaryPreferredLifetime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/temporary-valid-lifetime with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Autoconf{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Autoconf", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePath(t, md, goStruct)
	}
	return (&oc.QualifiedUint32{
		Metadata: md,
	}).SetVal(goStruct.GetTemporaryValidLifetime())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/temporary-valid-lifetime with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/temporary-valid-lifetime with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Autoconf{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Autoconf", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/temporary-valid-lifetime with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/temporary-valid-lifetime.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/temporary-valid-lifetime in the given batch object.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/temporary-valid-lifetime.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/temporary-valid-lifetime in the given batch object.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/temporary-valid-lifetime.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/autoconf/config/temporary-valid-lifetime in the given batch object.
func (n *Interface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePath extracts the value of the leaf TemporaryValidLifetime from its parent oc.Interface_Subinterface_Ipv6_Autoconf
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_Subinterface_Ipv6_Autoconf_TemporaryValidLifetimePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Autoconf) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.TemporaryValidLifetime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/dhcp-client with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_DhcpClientPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_DhcpClientPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetDhcpClient())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/dhcp-client with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_DhcpClientPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/dhcp-client with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_DhcpClientPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_DhcpClientPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/dhcp-client with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_DhcpClientPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/dhcp-client.
func (n *Interface_Subinterface_Ipv6_DhcpClientPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/dhcp-client in the given batch object.
func (n *Interface_Subinterface_Ipv6_DhcpClientPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/dhcp-client.
func (n *Interface_Subinterface_Ipv6_DhcpClientPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/dhcp-client in the given batch object.
func (n *Interface_Subinterface_Ipv6_DhcpClientPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/dhcp-client.
func (n *Interface_Subinterface_Ipv6_DhcpClientPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/dhcp-client in the given batch object.
func (n *Interface_Subinterface_Ipv6_DhcpClientPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv6_DhcpClientPath extracts the value of the leaf DhcpClient from its parent oc.Interface_Subinterface_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_Subinterface_Ipv6_DhcpClientPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.DhcpClient
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/dup-addr-detect-transmits with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_DupAddrDetectTransmitsPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_DupAddrDetectTransmitsPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint32{
		Metadata: md,
	}).SetVal(goStruct.GetDupAddrDetectTransmits())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/dup-addr-detect-transmits with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_DupAddrDetectTransmitsPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/dup-addr-detect-transmits with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_DupAddrDetectTransmitsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_DupAddrDetectTransmitsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/dup-addr-detect-transmits with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_DupAddrDetectTransmitsPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/dup-addr-detect-transmits.
func (n *Interface_Subinterface_Ipv6_DupAddrDetectTransmitsPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/dup-addr-detect-transmits in the given batch object.
func (n *Interface_Subinterface_Ipv6_DupAddrDetectTransmitsPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/dup-addr-detect-transmits.
func (n *Interface_Subinterface_Ipv6_DupAddrDetectTransmitsPath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/dup-addr-detect-transmits in the given batch object.
func (n *Interface_Subinterface_Ipv6_DupAddrDetectTransmitsPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/dup-addr-detect-transmits.
func (n *Interface_Subinterface_Ipv6_DupAddrDetectTransmitsPath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/dup-addr-detect-transmits in the given batch object.
func (n *Interface_Subinterface_Ipv6_DupAddrDetectTransmitsPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv6_DupAddrDetectTransmitsPath extracts the value of the leaf DupAddrDetectTransmits from its parent oc.Interface_Subinterface_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_Subinterface_Ipv6_DupAddrDetectTransmitsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.DupAddrDetectTransmits
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_EnabledPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnabled())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/enabled with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/enabled with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/enabled.
func (n *Interface_Subinterface_Ipv6_EnabledPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/enabled in the given batch object.
func (n *Interface_Subinterface_Ipv6_EnabledPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/enabled.
func (n *Interface_Subinterface_Ipv6_EnabledPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/enabled in the given batch object.
func (n *Interface_Subinterface_Ipv6_EnabledPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/enabled.
func (n *Interface_Subinterface_Ipv6_EnabledPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/enabled in the given batch object.
func (n *Interface_Subinterface_Ipv6_EnabledPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv6_EnabledPath extracts the value of the leaf Enabled from its parent oc.Interface_Subinterface_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_Subinterface_Ipv6_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Enabled
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/mtu with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_MtuPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_MtuPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/mtu with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_MtuPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/mtu with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_MtuPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_MtuPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/mtu with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_MtuPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/mtu.
func (n *Interface_Subinterface_Ipv6_MtuPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/mtu in the given batch object.
func (n *Interface_Subinterface_Ipv6_MtuPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/mtu.
func (n *Interface_Subinterface_Ipv6_MtuPath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/mtu in the given batch object.
func (n *Interface_Subinterface_Ipv6_MtuPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/mtu.
func (n *Interface_Subinterface_Ipv6_MtuPath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/config/mtu in the given batch object.
func (n *Interface_Subinterface_Ipv6_MtuPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv6_MtuPath extracts the value of the leaf Mtu from its parent oc.Interface_Subinterface_Ipv6
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_Subinterface_Ipv6_MtuPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Mtu
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_NeighborPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Ipv6_Neighbor {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Neighbor", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Ipv6_Neighbor{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_NeighborPath) Get(t testing.TB) *oc.Interface_Subinterface_Ipv6_Neighbor {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_NeighborPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Ipv6_Neighbor {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Ipv6_Neighbor
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Neighbor", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Ipv6_Neighbor{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_NeighborPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Ipv6_Neighbor {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Ipv6_Neighbor
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor.
func (n *Interface_Subinterface_Ipv6_NeighborPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor in the given batch object.
func (n *Interface_Subinterface_Ipv6_NeighborPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor.
func (n *Interface_Subinterface_Ipv6_NeighborPath) Replace(t testing.TB, val *oc.Interface_Subinterface_Ipv6_Neighbor) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor in the given batch object.
func (n *Interface_Subinterface_Ipv6_NeighborPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Ipv6_Neighbor) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor.
func (n *Interface_Subinterface_Ipv6_NeighborPath) Update(t testing.TB, val *oc.Interface_Subinterface_Ipv6_Neighbor) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor in the given batch object.
func (n *Interface_Subinterface_Ipv6_NeighborPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Ipv6_Neighbor) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/config/ip with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Neighbor_IpPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Neighbor", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_Neighbor_IpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/config/ip with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Neighbor_IpPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/config/ip with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Neighbor_IpPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Neighbor", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Neighbor_IpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/config/ip with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Neighbor_IpPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/config/ip.
func (n *Interface_Subinterface_Ipv6_Neighbor_IpPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/config/ip in the given batch object.
func (n *Interface_Subinterface_Ipv6_Neighbor_IpPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/config/ip.
func (n *Interface_Subinterface_Ipv6_Neighbor_IpPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/config/ip in the given batch object.
func (n *Interface_Subinterface_Ipv6_Neighbor_IpPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/config/ip.
func (n *Interface_Subinterface_Ipv6_Neighbor_IpPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/config/ip in the given batch object.
func (n *Interface_Subinterface_Ipv6_Neighbor_IpPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv6_Neighbor_IpPath extracts the value of the leaf Ip from its parent oc.Interface_Subinterface_Ipv6_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_Subinterface_Ipv6_Neighbor_IpPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Neighbor) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Ip
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/config/link-layer-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Neighbor", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/config/link-layer-address with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/config/link-layer-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Neighbor", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/config/link-layer-address with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/config/link-layer-address.
func (n *Interface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/config/link-layer-address in the given batch object.
func (n *Interface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/config/link-layer-address.
func (n *Interface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/config/link-layer-address in the given batch object.
func (n *Interface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/config/link-layer-address.
func (n *Interface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/neighbors/neighbor/config/link-layer-address in the given batch object.
func (n *Interface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPath extracts the value of the leaf LinkLayerAddress from its parent oc.Interface_Subinterface_Ipv6_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_Subinterface_Ipv6_Neighbor_LinkLayerAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Neighbor) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.LinkLayerAddress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisementPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_RouterAdvertisement{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_RouterAdvertisement", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisementPath) Get(t testing.TB) *oc.Interface_Subinterface_Ipv6_RouterAdvertisement {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisementPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_RouterAdvertisement{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_RouterAdvertisement", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Ipv6_RouterAdvertisement{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisementPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Ipv6_RouterAdvertisement {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Ipv6_RouterAdvertisement
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisementPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement in the given batch object.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisementPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisementPath) Replace(t testing.TB, val *oc.Interface_Subinterface_Ipv6_RouterAdvertisement) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement in the given batch object.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisementPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Ipv6_RouterAdvertisement) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisementPath) Update(t testing.TB, val *oc.Interface_Subinterface_Ipv6_RouterAdvertisement) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement in the given batch object.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisementPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Ipv6_RouterAdvertisement) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_IntervalPath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_RouterAdvertisement{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_RouterAdvertisement", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_RouterAdvertisement_IntervalPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/interval with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_IntervalPath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_IntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_RouterAdvertisement{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_RouterAdvertisement", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_RouterAdvertisement_IntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/interval with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_IntervalPathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/interval.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_IntervalPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/interval in the given batch object.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_IntervalPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/interval.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_IntervalPath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/interval in the given batch object.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_IntervalPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/interval.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_IntervalPath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/interval in the given batch object.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_IntervalPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv6_RouterAdvertisement_IntervalPath extracts the value of the leaf Interval from its parent oc.Interface_Subinterface_Ipv6_RouterAdvertisement
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_Subinterface_Ipv6_RouterAdvertisement_IntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_RouterAdvertisement) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Interval
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/lifetime with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_LifetimePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_RouterAdvertisement{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_RouterAdvertisement", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_RouterAdvertisement_LifetimePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/lifetime with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_LifetimePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/lifetime with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_LifetimePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_RouterAdvertisement{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_RouterAdvertisement", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_RouterAdvertisement_LifetimePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/lifetime with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_LifetimePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/lifetime.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_LifetimePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/lifetime in the given batch object.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_LifetimePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/lifetime.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_LifetimePath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/lifetime in the given batch object.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_LifetimePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/lifetime.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_LifetimePath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/lifetime in the given batch object.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_LifetimePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv6_RouterAdvertisement_LifetimePath extracts the value of the leaf Lifetime from its parent oc.Interface_Subinterface_Ipv6_RouterAdvertisement
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_Subinterface_Ipv6_RouterAdvertisement_LifetimePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_RouterAdvertisement) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Lifetime
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/suppress with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_SuppressPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_RouterAdvertisement{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_RouterAdvertisement", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_RouterAdvertisement_SuppressPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetSuppress())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/suppress with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_SuppressPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/suppress with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_SuppressPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_RouterAdvertisement{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_RouterAdvertisement", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_RouterAdvertisement_SuppressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/suppress with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_SuppressPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/suppress.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_SuppressPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/suppress in the given batch object.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_SuppressPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/suppress.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_SuppressPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/suppress in the given batch object.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_SuppressPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/suppress.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_SuppressPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/router-advertisement/config/suppress in the given batch object.
func (n *Interface_Subinterface_Ipv6_RouterAdvertisement_SuppressPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv6_RouterAdvertisement_SuppressPath extracts the value of the leaf Suppress from its parent oc.Interface_Subinterface_Ipv6_RouterAdvertisement
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_Subinterface_Ipv6_RouterAdvertisement_SuppressPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_RouterAdvertisement) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Suppress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_UnnumberedPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Unnumbered{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Unnumbered", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_UnnumberedPath) Get(t testing.TB) *oc.Interface_Subinterface_Ipv6_Unnumbered {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_UnnumberedPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Unnumbered{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Unnumbered", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_UnnumberedPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Ipv6_Unnumbered {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Ipv6_Unnumbered
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered.
func (n *Interface_Subinterface_Ipv6_UnnumberedPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered in the given batch object.
func (n *Interface_Subinterface_Ipv6_UnnumberedPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered.
func (n *Interface_Subinterface_Ipv6_UnnumberedPath) Replace(t testing.TB, val *oc.Interface_Subinterface_Ipv6_Unnumbered) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered in the given batch object.
func (n *Interface_Subinterface_Ipv6_UnnumberedPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Ipv6_Unnumbered) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered.
func (n *Interface_Subinterface_Ipv6_UnnumberedPath) Update(t testing.TB, val *oc.Interface_Subinterface_Ipv6_Unnumbered) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered in the given batch object.
func (n *Interface_Subinterface_Ipv6_UnnumberedPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Ipv6_Unnumbered) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/config/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Unnumbered_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Unnumbered{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Unnumbered", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_Unnumbered_EnabledPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnabled())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/config/enabled with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Unnumbered_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/config/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Unnumbered_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Unnumbered{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Unnumbered", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Unnumbered_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/config/enabled with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Unnumbered_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/config/enabled.
func (n *Interface_Subinterface_Ipv6_Unnumbered_EnabledPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/config/enabled in the given batch object.
func (n *Interface_Subinterface_Ipv6_Unnumbered_EnabledPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/config/enabled.
func (n *Interface_Subinterface_Ipv6_Unnumbered_EnabledPath) Replace(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/config/enabled in the given batch object.
func (n *Interface_Subinterface_Ipv6_Unnumbered_EnabledPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/config/enabled.
func (n *Interface_Subinterface_Ipv6_Unnumbered_EnabledPath) Update(t testing.TB, val bool) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/config/enabled in the given batch object.
func (n *Interface_Subinterface_Ipv6_Unnumbered_EnabledPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val bool) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv6_Unnumbered_EnabledPath extracts the value of the leaf Enabled from its parent oc.Interface_Subinterface_Ipv6_Unnumbered
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_Subinterface_Ipv6_Unnumbered_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Unnumbered) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Enabled
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefPath) Get(t testing.TB) *oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref in the given batch object.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefPath) Replace(t testing.TB, val *oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref in the given batch object.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefPath) Update(t testing.TB, val *oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref in the given batch object.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRefPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/config/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/config/interface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/config/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/config/interface with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/config/interface.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/config/interface in the given batch object.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/config/interface.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePath) Replace(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/config/interface in the given batch object.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/config/interface.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePath) Update(t testing.TB, val string) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/config/interface in the given batch object.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val string) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePath extracts the value of the leaf Interface from its parent oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef_InterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Interface
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/config/subinterface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/config/subinterface with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/config/subinterface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/config/subinterface with a ONCE subscription.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/config/subinterface.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/config/subinterface in the given batch object.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/config/subinterface.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath) Replace(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/config/subinterface in the given batch object.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/config/subinterface.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath) Update(t testing.TB, val uint32) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/ipv6/unnumbered/interface-ref/config/subinterface in the given batch object.
func (n *Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint32) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath extracts the value of the leaf Subinterface from its parent oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_Subinterface_Ipv6_Unnumbered_InterfaceRef_SubinterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Ipv6_Unnumbered_InterfaceRef) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Subinterface
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_VlanPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Vlan{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_VlanPath) Get(t testing.TB) *oc.Interface_Subinterface_Vlan {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_VlanPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Vlan{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan with a ONCE subscription.
func (n *Interface_Subinterface_VlanPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Vlan {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Vlan
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan.
func (n *Interface_Subinterface_VlanPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan in the given batch object.
func (n *Interface_Subinterface_VlanPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan.
func (n *Interface_Subinterface_VlanPath) Replace(t testing.TB, val *oc.Interface_Subinterface_Vlan) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan in the given batch object.
func (n *Interface_Subinterface_VlanPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Vlan) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan.
func (n *Interface_Subinterface_VlanPath) Update(t testing.TB, val *oc.Interface_Subinterface_Vlan) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan in the given batch object.
func (n *Interface_Subinterface_VlanPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Vlan) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_EgressMappingPath) Lookup(t testing.TB) *oc.QualifiedInterface_Subinterface_Vlan_EgressMapping {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_EgressMapping{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_EgressMapping", goStruct, false, true)
	if ok {
		return (&oc.QualifiedInterface_Subinterface_Vlan_EgressMapping{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_EgressMappingPath) Get(t testing.TB) *oc.Interface_Subinterface_Vlan_EgressMapping {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_EgressMappingPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_Subinterface_Vlan_EgressMapping {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_Subinterface_Vlan_EgressMapping
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_EgressMapping{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_EgressMapping", goStruct, queryPath, false, true)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_Subinterface_Vlan_EgressMapping{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_EgressMappingPathAny) Get(t testing.TB) []*oc.Interface_Subinterface_Vlan_EgressMapping {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_Subinterface_Vlan_EgressMapping
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping.
func (n *Interface_Subinterface_Vlan_EgressMappingPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping in the given batch object.
func (n *Interface_Subinterface_Vlan_EgressMappingPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping.
func (n *Interface_Subinterface_Vlan_EgressMappingPath) Replace(t testing.TB, val *oc.Interface_Subinterface_Vlan_EgressMapping) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping in the given batch object.
func (n *Interface_Subinterface_Vlan_EgressMappingPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Vlan_EgressMapping) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping.
func (n *Interface_Subinterface_Vlan_EgressMappingPath) Update(t testing.TB, val *oc.Interface_Subinterface_Vlan_EgressMapping) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping in the given batch object.
func (n *Interface_Subinterface_Vlan_EgressMappingPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val *oc.Interface_Subinterface_Vlan_EgressMapping) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/tpid with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPath) Lookup(t testing.TB) *oc.QualifiedE_VlanTypes_TPID_TYPES {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_EgressMapping{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_EgressMapping", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_EgressMapping_TpidPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/tpid with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPath) Get(t testing.TB) oc.E_VlanTypes_TPID_TYPES {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/tpid with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPathAny) Lookup(t testing.TB) []*oc.QualifiedE_VlanTypes_TPID_TYPES {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_VlanTypes_TPID_TYPES
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_EgressMapping{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_EgressMapping", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_EgressMapping_TpidPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/tpid with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPathAny) Get(t testing.TB) []oc.E_VlanTypes_TPID_TYPES {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_VlanTypes_TPID_TYPES
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/tpid.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/tpid in the given batch object.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/tpid.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPath) Replace(t testing.TB, val oc.E_VlanTypes_TPID_TYPES) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/tpid in the given batch object.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val oc.E_VlanTypes_TPID_TYPES) {
	t.Helper()
	b.BatchReplace(t, n, val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/tpid.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPath) Update(t testing.TB, val oc.E_VlanTypes_TPID_TYPES) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/tpid in the given batch object.
func (n *Interface_Subinterface_Vlan_EgressMapping_TpidPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val oc.E_VlanTypes_TPID_TYPES) {
	t.Helper()
	b.BatchUpdate(t, n, val)
}

// convertInterface_Subinterface_Vlan_EgressMapping_TpidPath extracts the value of the leaf Tpid from its parent oc.Interface_Subinterface_Vlan_EgressMapping
// and combines the update with an existing Metadata to return a *oc.QualifiedE_VlanTypes_TPID_TYPES.
func convertInterface_Subinterface_Vlan_EgressMapping_TpidPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_EgressMapping) *oc.QualifiedE_VlanTypes_TPID_TYPES {
	t.Helper()
	qv := &oc.QualifiedE_VlanTypes_TPID_TYPES{
		Metadata: md,
	}
	val := parent.Tpid
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-id with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_Subinterface_Vlan_EgressMapping{}
	md, ok := oc.Lookup(t, n, "Interface_Subinterface_Vlan_EgressMapping", goStruct, true, true)
	if ok {
		return convertInterface_Subinterface_Vlan_EgressMapping_VlanIdPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-id with a ONCE subscription,
// failing the test fatally if no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-id with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_Subinterface_Vlan_EgressMapping{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_Subinterface_Vlan_EgressMapping", goStruct, queryPath, true, true)
		if !ok {
			continue
		}
		qv := convertInterface_Subinterface_Vlan_EgressMapping_VlanIdPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-id with a ONCE subscription.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Delete deletes the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-id.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPath) Delete(t testing.TB) *gpb.SetResponse {
	t.Helper()
	return genutil.Delete(t, n)
}

// BatchDelete buffers a config delete operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPath) BatchDelete(t testing.TB, b *config.SetRequestBatch) {
	t.Helper()
	b.BatchDelete(t, n)
}

// Replace replaces the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-id.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPath) Replace(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Replace(t, n, &val)
}

// BatchReplace buffers a config replace operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPath) BatchReplace(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchReplace(t, n, &val)
}

// Update updates the configuration at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-id.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPath) Update(t testing.TB, val uint16) *gpb.SetResponse {
	t.Helper()
	return genutil.Update(t, n, &val)
}

// BatchUpdate buffers a config update operation at /openconfig-interfaces/interfaces/interface/subinterfaces/subinterface/vlan/egress-mapping/config/vlan-id in the given batch object.
func (n *Interface_Subinterface_Vlan_EgressMapping_VlanIdPath) BatchUpdate(t testing.TB, b *config.SetRequestBatch, val uint16) {
	t.Helper()
	b.BatchUpdate(t, n, &val)
}

// convertInterface_Subinterface_Vlan_EgressMapping_VlanIdPath extracts the value of the leaf VlanId from its parent oc.Interface_Subinterface_Vlan_EgressMapping
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_Subinterface_Vlan_EgressMapping_VlanIdPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_Subinterface_Vlan_EgressMapping) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.VlanId
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
