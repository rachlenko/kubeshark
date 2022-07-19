// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64
// +build 386 amd64

package tlstapper

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type tlsTapper46GoidOffsets struct {
	G_addrOffset uint64
	GoidOffset   uint64
}

type tlsTapper46TlsChunk struct {
	Pid         uint32
	Tgid        uint32
	Len         uint32
	Start       uint32
	Recorded    uint32
	Fd          uint32
	Flags       uint32
	AddressInfo struct {
		Mode  int32
		Saddr uint32
		Daddr uint32
		Sport uint16
		Dport uint16
	}
	Data [4096]uint8
}

// loadTlsTapper46 returns the embedded CollectionSpec for tlsTapper46.
func loadTlsTapper46() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_TlsTapper46Bytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load tlsTapper46: %w", err)
	}

	return spec, err
}

// loadTlsTapper46Objects loads tlsTapper46 and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//     *tlsTapper46Objects
//     *tlsTapper46Programs
//     *tlsTapper46Maps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadTlsTapper46Objects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadTlsTapper46()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// tlsTapper46Specs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type tlsTapper46Specs struct {
	tlsTapper46ProgramSpecs
	tlsTapper46MapSpecs
}

// tlsTapper46Specs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type tlsTapper46ProgramSpecs struct {
	GoCryptoTlsAbi0Read           *ebpf.ProgramSpec `ebpf:"go_crypto_tls_abi0_read"`
	GoCryptoTlsAbi0ReadEx         *ebpf.ProgramSpec `ebpf:"go_crypto_tls_abi0_read_ex"`
	GoCryptoTlsAbi0Write          *ebpf.ProgramSpec `ebpf:"go_crypto_tls_abi0_write"`
	GoCryptoTlsAbi0WriteEx        *ebpf.ProgramSpec `ebpf:"go_crypto_tls_abi0_write_ex"`
	GoCryptoTlsAbiInternalRead    *ebpf.ProgramSpec `ebpf:"go_crypto_tls_abi_internal_read"`
	GoCryptoTlsAbiInternalReadEx  *ebpf.ProgramSpec `ebpf:"go_crypto_tls_abi_internal_read_ex"`
	GoCryptoTlsAbiInternalWrite   *ebpf.ProgramSpec `ebpf:"go_crypto_tls_abi_internal_write"`
	GoCryptoTlsAbiInternalWriteEx *ebpf.ProgramSpec `ebpf:"go_crypto_tls_abi_internal_write_ex"`
	SslRead                       *ebpf.ProgramSpec `ebpf:"ssl_read"`
	SslReadEx                     *ebpf.ProgramSpec `ebpf:"ssl_read_ex"`
	SslRetRead                    *ebpf.ProgramSpec `ebpf:"ssl_ret_read"`
	SslRetReadEx                  *ebpf.ProgramSpec `ebpf:"ssl_ret_read_ex"`
	SslRetWrite                   *ebpf.ProgramSpec `ebpf:"ssl_ret_write"`
	SslRetWriteEx                 *ebpf.ProgramSpec `ebpf:"ssl_ret_write_ex"`
	SslWrite                      *ebpf.ProgramSpec `ebpf:"ssl_write"`
	SslWriteEx                    *ebpf.ProgramSpec `ebpf:"ssl_write_ex"`
	SysEnterAccept4               *ebpf.ProgramSpec `ebpf:"sys_enter_accept4"`
	SysEnterConnect               *ebpf.ProgramSpec `ebpf:"sys_enter_connect"`
	SysEnterRead                  *ebpf.ProgramSpec `ebpf:"sys_enter_read"`
	SysEnterWrite                 *ebpf.ProgramSpec `ebpf:"sys_enter_write"`
	SysExitAccept4                *ebpf.ProgramSpec `ebpf:"sys_exit_accept4"`
	SysExitConnect                *ebpf.ProgramSpec `ebpf:"sys_exit_connect"`
	TcpRecvmsg                    *ebpf.ProgramSpec `ebpf:"tcp_recvmsg"`
	TcpSendmsg                    *ebpf.ProgramSpec `ebpf:"tcp_sendmsg"`
}

// tlsTapper46MapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type tlsTapper46MapSpecs struct {
	AcceptSyscallContext *ebpf.MapSpec `ebpf:"accept_syscall_context"`
	ChunksBuffer         *ebpf.MapSpec `ebpf:"chunks_buffer"`
	ConnectSyscallInfo   *ebpf.MapSpec `ebpf:"connect_syscall_info"`
	FileDescriptorToIpv4 *ebpf.MapSpec `ebpf:"file_descriptor_to_ipv4"`
	GoReadContext        *ebpf.MapSpec `ebpf:"go_read_context"`
	GoWriteContext       *ebpf.MapSpec `ebpf:"go_write_context"`
	GoidOffsetsMap       *ebpf.MapSpec `ebpf:"goid_offsets_map"`
	Heap                 *ebpf.MapSpec `ebpf:"heap"`
	LogBuffer            *ebpf.MapSpec `ebpf:"log_buffer"`
	OpensslReadContext   *ebpf.MapSpec `ebpf:"openssl_read_context"`
	OpensslWriteContext  *ebpf.MapSpec `ebpf:"openssl_write_context"`
	PidsMap              *ebpf.MapSpec `ebpf:"pids_map"`
}

// tlsTapper46Objects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadTlsTapper46Objects or ebpf.CollectionSpec.LoadAndAssign.
type tlsTapper46Objects struct {
	tlsTapper46Programs
	tlsTapper46Maps
}

func (o *tlsTapper46Objects) Close() error {
	return _TlsTapper46Close(
		&o.tlsTapper46Programs,
		&o.tlsTapper46Maps,
	)
}

// tlsTapper46Maps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadTlsTapper46Objects or ebpf.CollectionSpec.LoadAndAssign.
type tlsTapper46Maps struct {
	AcceptSyscallContext *ebpf.Map `ebpf:"accept_syscall_context"`
	ChunksBuffer         *ebpf.Map `ebpf:"chunks_buffer"`
	ConnectSyscallInfo   *ebpf.Map `ebpf:"connect_syscall_info"`
	FileDescriptorToIpv4 *ebpf.Map `ebpf:"file_descriptor_to_ipv4"`
	GoReadContext        *ebpf.Map `ebpf:"go_read_context"`
	GoWriteContext       *ebpf.Map `ebpf:"go_write_context"`
	GoidOffsetsMap       *ebpf.Map `ebpf:"goid_offsets_map"`
	Heap                 *ebpf.Map `ebpf:"heap"`
	LogBuffer            *ebpf.Map `ebpf:"log_buffer"`
	OpensslReadContext   *ebpf.Map `ebpf:"openssl_read_context"`
	OpensslWriteContext  *ebpf.Map `ebpf:"openssl_write_context"`
	PidsMap              *ebpf.Map `ebpf:"pids_map"`
}

func (m *tlsTapper46Maps) Close() error {
	return _TlsTapper46Close(
		m.AcceptSyscallContext,
		m.ChunksBuffer,
		m.ConnectSyscallInfo,
		m.FileDescriptorToIpv4,
		m.GoReadContext,
		m.GoWriteContext,
		m.GoidOffsetsMap,
		m.Heap,
		m.LogBuffer,
		m.OpensslReadContext,
		m.OpensslWriteContext,
		m.PidsMap,
	)
}

// tlsTapper46Programs contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadTlsTapper46Objects or ebpf.CollectionSpec.LoadAndAssign.
type tlsTapper46Programs struct {
	GoCryptoTlsAbi0Read           *ebpf.Program `ebpf:"go_crypto_tls_abi0_read"`
	GoCryptoTlsAbi0ReadEx         *ebpf.Program `ebpf:"go_crypto_tls_abi0_read_ex"`
	GoCryptoTlsAbi0Write          *ebpf.Program `ebpf:"go_crypto_tls_abi0_write"`
	GoCryptoTlsAbi0WriteEx        *ebpf.Program `ebpf:"go_crypto_tls_abi0_write_ex"`
	GoCryptoTlsAbiInternalRead    *ebpf.Program `ebpf:"go_crypto_tls_abi_internal_read"`
	GoCryptoTlsAbiInternalReadEx  *ebpf.Program `ebpf:"go_crypto_tls_abi_internal_read_ex"`
	GoCryptoTlsAbiInternalWrite   *ebpf.Program `ebpf:"go_crypto_tls_abi_internal_write"`
	GoCryptoTlsAbiInternalWriteEx *ebpf.Program `ebpf:"go_crypto_tls_abi_internal_write_ex"`
	SslRead                       *ebpf.Program `ebpf:"ssl_read"`
	SslReadEx                     *ebpf.Program `ebpf:"ssl_read_ex"`
	SslRetRead                    *ebpf.Program `ebpf:"ssl_ret_read"`
	SslRetReadEx                  *ebpf.Program `ebpf:"ssl_ret_read_ex"`
	SslRetWrite                   *ebpf.Program `ebpf:"ssl_ret_write"`
	SslRetWriteEx                 *ebpf.Program `ebpf:"ssl_ret_write_ex"`
	SslWrite                      *ebpf.Program `ebpf:"ssl_write"`
	SslWriteEx                    *ebpf.Program `ebpf:"ssl_write_ex"`
	SysEnterAccept4               *ebpf.Program `ebpf:"sys_enter_accept4"`
	SysEnterConnect               *ebpf.Program `ebpf:"sys_enter_connect"`
	SysEnterRead                  *ebpf.Program `ebpf:"sys_enter_read"`
	SysEnterWrite                 *ebpf.Program `ebpf:"sys_enter_write"`
	SysExitAccept4                *ebpf.Program `ebpf:"sys_exit_accept4"`
	SysExitConnect                *ebpf.Program `ebpf:"sys_exit_connect"`
	TcpRecvmsg                    *ebpf.Program `ebpf:"tcp_recvmsg"`
	TcpSendmsg                    *ebpf.Program `ebpf:"tcp_sendmsg"`
}

func (p *tlsTapper46Programs) Close() error {
	return _TlsTapper46Close(
		p.GoCryptoTlsAbi0Read,
		p.GoCryptoTlsAbi0ReadEx,
		p.GoCryptoTlsAbi0Write,
		p.GoCryptoTlsAbi0WriteEx,
		p.GoCryptoTlsAbiInternalRead,
		p.GoCryptoTlsAbiInternalReadEx,
		p.GoCryptoTlsAbiInternalWrite,
		p.GoCryptoTlsAbiInternalWriteEx,
		p.SslRead,
		p.SslReadEx,
		p.SslRetRead,
		p.SslRetReadEx,
		p.SslRetWrite,
		p.SslRetWriteEx,
		p.SslWrite,
		p.SslWriteEx,
		p.SysEnterAccept4,
		p.SysEnterConnect,
		p.SysEnterRead,
		p.SysEnterWrite,
		p.SysExitAccept4,
		p.SysExitConnect,
		p.TcpRecvmsg,
		p.TcpSendmsg,
	)
}

func _TlsTapper46Close(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//go:embed tlstapper46_bpfel_x86.o
var _TlsTapper46Bytes []byte