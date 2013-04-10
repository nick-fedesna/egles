// +build ignore
package drm

//#include <linux/types.h>
//#include <stddef.h>
//#include <asm/ioctl.h>
//#include <libdrm/drm.h>
//#include <libdrm/drm_mode.h>
import "C"

type (
	Handle            C.drm_handle_t
	Context           C.drm_context_t
	Drawable          C.drm_drawable_t
	Magic             C.drm_magic_t
	ClipRect          C.struct_drm_clip_rect
	DrawableInfo      C.struct_drm_drawable_info
	TexRegion         C.struct_drm_tex_region
	HwLock            C.struct_drm_hw_lock
	Version           C.struct_drm_version
	Unique            C.struct_drm_unique
	List              C.struct_drm_list
	Block             C.struct_drm_block
	Control           C.struct_drm_control
	MapType           C.enum_drm_map_type
	MapFlags          C.enum_drm_map_flags
	CtxPrivmap        C.struct_ctx_priv_map
	Map               C.struct_drm_map
	Client            C.struct_drm_client
	StatType          C.enum_drm_stat_type
	Stats             C.struct_drm_stats
	LockFlags         C.enum_drm_lock_flags
	Lock              C.struct_drm_lock
	DmaFlags          C.enum_drm_dma_flags
	BufDesk           C.struct_drm_buf_desc
	BufInfo           C.struct_drm_buf_info
	BufFree           C.struct_drm_buf_free
	BufPub            C.struct_drm_buf_pub
	BufMap            C.struct_drm_buf_map
	Dma               C.struct_drm_dma
	CtxFlags          C.enum_drm_ctx_flags
	Ctx               C.struct_drm_ctx
	CtxRes            C.struct_drm_ctx_res
	Draw              C.struct_drm_draw
	DrawableInfoType  C.drm_drawable_info_type_t
	UpdateDraw        C.struct_drm_update_draw
	Auth              C.struct_drm_auth
	IrqBusid          C.struct_drm_irq_busid
	VBlankSeqType     C.enum_drm_vblank_seq_type
	WaitVBlankRequest C.struct_drm_wait_vblank_request
	WaitVBlankReply   C.struct_drm_wait_vblank_reply
	WaitVBlank        C.union_drm_wait_vblank
	ModesetCtl        C.struct_drm_modeset_ctl
	AgpMode           C.struct_drm_agp_mode
	AgpBuffer         C.struct_drm_agp_buffer
	AgpBinding        C.struct_drm_agp_binding
	AgpInfo           C.struct_drm_agp_info
	ScatterGather     C.struct_drm_scatter_gather
	SetVersion        C.struct_drm_set_version
	GemClose          C.struct_drm_gem_close
	GemFlink          C.struct_drm_gem_flink
	GemOpen           C.struct_drm_gem_open
	GetCap            C.struct_drm_get_cap
	PrimeHandle       C.struct_drm_prime_handle
	Event             C.struct_drm_event
	EventVBlank       C.struct_drm_event_vblank
)

const (
	NAME        = C.DRM_NAME
	MIN_ORDER   = C.DRM_MIN_ORDER
	MAX_ORDER   = C.DRM_MAX_ORDER
	RAM_PERCENT = C.DRM_RAM_PERCENT
	LOCK_HELD   = C._DRM_LOCK_HELD
	LOCK_CONT   = C._DRM_LOCK_CONT

	//drm_map_type
	FRAME_BUFFER   = C._DRM_FRAME_BUFFER
	REGISTERS      = C._DRM_REGISTERS
	SHM            = C._DRM_SHM
	AGP            = C._DRM_AGP
	SCATTER_GATHER = C._DRM_SCATTER_GATHER
	CONSISTENT     = C._DRM_CONSISTENT
	GEM            = C._DRM_GEM

	//drm_map_flags
	RESTRICTED      = C._DRM_RESTRICTED
	READ_ONLY       = C._DRM_READ_ONLY
	LOCKED          = C._DRM_LOCKED
	KERNEL          = C._DRM_KERNEL
	WRITE_COMBINING = C._DRM_WRITE_COMBINING
	CONTAINS_LOCK   = C._DRM_CONTAINS_LOCK
	REMOVABLE       = C._DRM_REMOVABLE
	DRIVER          = C._DRM_DRIVER
	//drm_stat_type
	STAT_LOCK      = C._DRM_STAT_LOCK
	STAT_OPENS     = C._DRM_STAT_OPENS
	STAT_CLOSES    = C._DRM_STAT_CLOSES
	STAT_IOCTLS    = C._DRM_STAT_IOCTLS
	STAT_LOCKS     = C._DRM_STAT_LOCKS
	STAT_UNLOCKS   = C._DRM_STAT_UNLOCKS
	STAT_VALUE     = C._DRM_STAT_VALUE
	STAT_BYTE      = C._DRM_STAT_BYTE
	STAT_COUNT     = C._DRM_STAT_COUNT
	STAT_IRQ       = C._DRM_STAT_IRQ
	STAT_PRIMARY   = C._DRM_STAT_PRIMARY
	STAT_SECONDARY = C._DRM_STAT_SECONDARY
	STAT_DMA       = C._DRM_STAT_DMA
	STAT_SPECIAL   = C._DRM_STAT_SPECIAL
	STAT_MISSED    = C._DRM_STAT_MISSED
	//drm_lock_flags
	LOCK_READY      = C._DRM_LOCK_READY
	LOCK_QUIESCENT  = C._DRM_LOCK_QUIESCENT
	LOCK_FLUSH      = C._DRM_LOCK_FLUSH
	LOCK_FLUSH_ALL  = C._DRM_LOCK_FLUSH_ALL
	HALT_ALL_QUEUES = C._DRM_HALT_ALL_QUEUES
	HALT_CUR_QUEUES = C._DRM_HALT_CUR_QUEUES
	//drm_dma_flags
	DMA_BLOCK        = C._DRM_DMA_BLOCK
	DMA_WHILE_LOCKED = C._DRM_DMA_WHILE_LOCKED
	DMA_PRIORITY     = C._DRM_DMA_PRIORITY
	DMA_WaIT         = C._DRM_DMA_WAIT
	DMA_SMALLER_OK   = C._DRM_DMA_SMALLER_OK
	DMA_LARGER_OK    = C._DRM_DMA_LARGER_OK
	//drm_buf_desc.flags
	PAGE_ALIGN    = C._DRM_PAGE_ALIGN
	AGP_BUFFER    = C._DRM_AGP_BUFFER
	SG_BUFFER     = C._DRM_SG_BUFFER
	FB_BUFFER     = C._DRM_FB_BUFFER
	PCI_BUFFER_RO = C._DRM_PCI_BUFFER_RO
	//drm_ctx_flags
	CONTEXT_PRESERVED = C._DRM_CONTEXT_PRESERVED
	CONTEXT_2DONLY    = C._DRM_CONTEXT_2DONLY
	//drm_vblank_seq_type
	VBLANK_ABSOLUTE   = C._DRM_VBLANK_ABSOLUTE
	VBLANK_RELATIVE   = C._DRM_VBLANK_RELATIVE
	VBLANK_EVENT      = C._DRM_VBLANK_EVENT
	VBLANK_FLIP       = C._DRM_VBLANK_FLIP
	VBLANK_NEXTONMISS = C._DRM_VBLANK_NEXTONMISS
	VBLANK_SECONDARY  = C._DRM_VBLANK_SECONDARY
	VBLANK_SIGNAL     = C._DRM_VBLANK_SIGNAL

	VBLANK_TYPES_MASK = C._DRM_VBLANK_TYPES_MASK
	VBLANK_FLAGS_MASK = C._DRM_VBLANK_FLAGS_MASK
	PRE_MODESET       = C._DRM_PRE_MODESET
	POST_MODESET      = C._DRM_POST_MODESET
	//CLOEXEC           = C.DRM_CLOEXEC
)

//drm_mode.h
const (
	DISPLAY_INFO_LEN   = C.DRM_DISPLAY_INFO_LEN
	CONNECTOR_NAME_LEN = C.DRM_CONNECTOR_NAME_LEN
	DISPLAY_MODE_LEN   = C.DRM_DISPLAY_MODE_LEN
	PROP_NAME_LEN      = C.DRM_PROP_NAME_LEN

	MODE_TYPE_BUILTIN   = C.DRM_MODE_TYPE_BUILTIN
	MODE_TYPE_CLOCK_C   = C.DRM_MODE_TYPE_CLOCK_C
	MODE_TYPE_CRTC_C    = C.DRM_MODE_TYPE_CRTC_C
	MODE_TYPE_PREFERRED = C.DRM_MODE_TYPE_PREFERRED
	MODE_TYPE_DEFAULT   = C.DRM_MODE_TYPE_DEFAULT
	MODE_TYPE_USERDEF   = C.DRM_MODE_TYPE_USERDEF
	MODE_TYPE_DRIVER    = C.DRM_MODE_TYPE_DRIVER

	/* Video mode flags */

	// bit compatible with the xorg definitions. 
	MODE_FLAG_PHSYNC    = C.DRM_MODE_FLAG_PHSYNC
	MODE_FLAG_NHSYNC    = C.DRM_MODE_FLAG_NHSYNC
	MODE_FLAG_PVSYNC    = C.DRM_MODE_FLAG_PVSYNC
	MODE_FLAG_NVSYNC    = C.DRM_MODE_FLAG_NVSYNC
	MODE_FLAG_INTERLACE = C.DRM_MODE_FLAG_INTERLACE
	MODE_FLAG_DBLSCAN   = C.DRM_MODE_FLAG_DBLSCAN
	MODE_FLAG_CSYNC     = C.DRM_MODE_FLAG_CSYNC
	MODE_FLAG_PCSYNC    = C.DRM_MODE_FLAG_PCSYNC
	MODE_FLAG_NCSYNC    = C.DRM_MODE_FLAG_NCSYNC

	// hskew provided 
	MODE_FLAG_HSKEW   = C.DRM_MODE_FLAG_HSKEW
	MODE_FLAG_BCAST   = C.DRM_MODE_FLAG_BCAST
	MODE_FLAG_PIXMUX  = C.DRM_MODE_FLAG_PIXMUX
	MODE_FLAG_DBLCLK  = C.DRM_MODE_FLAG_DBLCLK
	MODE_FLAG_CLKDIV2 = C.DRM_MODE_FLAG_CLKDIV2

	/* DPMS flags */
	// bit compatible with the xorg definitions. 
	MODE_DPMS_ON      = C.DRM_MODE_DPMS_ON
	MODE_DPMS_STANDBY = C.DRM_MODE_DPMS_STANDBY
	MODE_DPMS_SUSPEND = C.DRM_MODE_DPMS_SUSPEND
	MODE_DPMS_OFF     = C.DRM_MODE_DPMS_OFF

	/* Scaling mode options */
	// Unmodified timing (display or software can still scale) 
	MODE_SCALE_NONE = C.DRM_MODE_SCALE_NONE
	// Full screen, ignore aspect 
	MODE_SCALE_FULLSCREEN = C.DRM_MODE_SCALE_FULLSCREEN
	// Centered, no scaling 
	MODE_SCALE_CENTER = C.DRM_MODE_SCALE_CENTER
	// Full screen, preserve aspect 
	MODE_SCALE_ASPECT = C.DRM_MODE_SCALE_ASPECT

	/* Dithering mode options */
	MODE_DITHERING_OFF  = C.DRM_MODE_DITHERING_OFF
	MODE_DITHERING_ON   = C.DRM_MODE_DITHERING_ON
	MODE_DITHERING_AUTO = C.DRM_MODE_DITHERING_AUTO

	/* Dirty info options */
	MODE_DIRTY_OFF      = C.DRM_MODE_DIRTY_OFF
	MODE_DIRTY_ON       = C.DRM_MODE_DIRTY_ON
	MODE_DIRTY_ANNOTATE = C.DRM_MODE_DIRTY_ANNOTATE
)

type (
	ModeModeInfo C.struct_drm_mode_modeinfo
	ModeCardRes  C.struct_drm_mode_card_res
	ModeCrtc     C.struct_drm_mode_crtc
)

const (
	MODE_PRESENT_TOP_FIELD    = C.DRM_MODE_PRESENT_TOP_FIELD
	MODE_PRESENT_BOTTOM_FIELD = C.DRM_MODE_PRESENT_BOTTOM_FIELD
)

type (
	/* Planes blend with or override other bits on the CRTC */
	ModeSetPlane    C.struct_drm_mode_set_plane
	ModeGetPlane    C.struct_drm_mode_get_plane
	ModeGetPlaneRes C.struct_drm_mode_get_plane_res
)

const (
	MODE_ENCODER_NONE  = C.DRM_MODE_ENCODER_NONE
	MODE_ENCODER_DAC   = C.DRM_MODE_ENCODER_DAC
	MODE_ENCODER_TMDS  = C.DRM_MODE_ENCODER_TMDS
	MODE_ENCODER_LVDS  = C.DRM_MODE_ENCODER_LVDS
	MODE_ENCODER_TVDAC = C.DRM_MODE_ENCODER_TVDAC
)

type ModeGetEncoder C.struct_drm_mode_get_encoder

/* 
This is for connectors with multiple signal types. 
Try to match DRM_MODE_CONNECTOR_X as closely as possible. 
*/
const (
	MODE_SUBCONNECTOR_Automatic = C.DRM_MODE_SUBCONNECTOR_Automatic
	MODE_SUBCONNECTOR_Unknown   = C.DRM_MODE_SUBCONNECTOR_Unknown
	MODE_SUBCONNECTOR_DVID      = C.DRM_MODE_SUBCONNECTOR_DVID
	MODE_SUBCONNECTOR_DVIA      = C.DRM_MODE_SUBCONNECTOR_DVIA
	MODE_SUBCONNECTOR_Composite = C.DRM_MODE_SUBCONNECTOR_Composite
	MODE_SUBCONNECTOR_SVIDEO    = C.DRM_MODE_SUBCONNECTOR_SVIDEO
	MODE_SUBCONNECTOR_Component = C.DRM_MODE_SUBCONNECTOR_Component
	MODE_SUBCONNECTOR_SCART     = C.DRM_MODE_SUBCONNECTOR_SCART

	MODE_CONNECTOR_Unknown     = C.DRM_MODE_CONNECTOR_Unknown
	MODE_CONNECTOR_VGA         = C.DRM_MODE_CONNECTOR_VGA
	MODE_CONNECTOR_DVII        = C.DRM_MODE_CONNECTOR_DVII
	MODE_CONNECTOR_DVID        = C.DRM_MODE_CONNECTOR_DVID
	MODE_CONNECTOR_DVIA        = C.DRM_MODE_CONNECTOR_DVIA
	MODE_CONNECTOR_Composite   = C.DRM_MODE_CONNECTOR_Composite
	MODE_CONNECTOR_SVIDEO      = C.DRM_MODE_CONNECTOR_SVIDEO
	MODE_CONNECTOR_LVDS        = C.DRM_MODE_CONNECTOR_LVDS
	MODE_CONNECTOR_Component   = C.DRM_MODE_CONNECTOR_Component
	MODE_CONNECTOR_9PinDIN     = C.DRM_MODE_CONNECTOR_9PinDIN
	MODE_CONNECTOR_DisplayPort = C.DRM_MODE_CONNECTOR_DisplayPort
	MODE_CONNECTOR_HDMIA       = C.DRM_MODE_CONNECTOR_HDMIA
	MODE_CONNECTOR_HDMIB       = C.DRM_MODE_CONNECTOR_HDMIB
	MODE_CONNECTOR_TV          = C.DRM_MODE_CONNECTOR_TV
	MODE_CONNECTOR_eDP         = C.DRM_MODE_CONNECTOR_eDP
)

type ModeGetConnector C.struct_drm_mode_get_connector

const (
	MODE_PROP_PENDING   = C.DRM_MODE_PROP_PENDING
	MODE_PROP_RANGE     = C.DRM_MODE_PROP_RANGE
	MODE_PROP_IMMUTABLE = C.DRM_MODE_PROP_IMMUTABLE
	// enumerated type with text strings */
	MODE_PROP_ENUM = C.DRM_MODE_PROP_ENUM
	MODE_PROP_BLOB = C.DRM_MODE_PROP_BLOB
	// bitmask of enumerated types */
	MODE_PROP_BITMASK = C.DRM_MODE_PROP_BITMASK
)

type (
	ModePropertyEnum         C.struct_drm_mode_property_enum
	ModeGetProperty          C.struct_drm_mode_get_property
	ModeConnectorSetProperty C.struct_drm_mode_connector_set_property
)

const (
	MODE_OBJECT_CRTC      = C.DRM_MODE_OBJECT_CRTC
	MODE_OBJECT_CONNECTOR = C.DRM_MODE_OBJECT_CONNECTOR
	MODE_OBJECT_ENCODER   = C.DRM_MODE_OBJECT_ENCODER
	MODE_OBJECT_MODE      = C.DRM_MODE_OBJECT_MODE
	MODE_OBJECT_PROPERTY  = C.DRM_MODE_OBJECT_PROPERTY
	MODE_OBJECT_FB        = C.DRM_MODE_OBJECT_FB
	MODE_OBJECT_BLOB      = C.DRM_MODE_OBJECT_BLOB
	MODE_OBJECT_PLANE     = C.DRM_MODE_OBJECT_PLANE
)

type (
	ModeObjGetProperties C.struct_drm_mode_obj_get_properties

	ModeObjSetProperty C.struct_drm_mode_obj_set_property

	ModeGetBlob C.struct_drm_mode_get_blob

	ModeFbCmd C.struct_drm_mode_fb_cmd
)

// for interlaced framebuffers */
const MODE_FB_INTERLACED = C.DRM_MODE_FB_INTERLACED

type ModeFbCmd2 C.struct_drm_mode_fb_cmd2

const (
	MODE_FB_DIRTY_ANNOTATE_COPY = C.DRM_MODE_FB_DIRTY_ANNOTATE_COPY
	MODE_FB_DIRTY_ANNOTATE_FILL = C.DRM_MODE_FB_DIRTY_ANNOTATE_FILL
	MODE_FB_DIRTY_FLAGS         = C.DRM_MODE_FB_DIRTY_FLAGS
)

type ModeFbDirtyCmd C.struct_drm_mode_fb_dirty_cmd

type ModeModeCmd C.struct_drm_mode_mode_cmd

const (
	MODE_CURSOR_BO   = C.DRM_MODE_CURSOR_BO
	MODE_CURSOR_MOVE = C.DRM_MODE_CURSOR_MOVE
)

/* 
depending on the value in flags diffrent members are used.
CURSOR_BO uses crtc width height handle - if 0 turns the cursor of
CURSOR_MOVE uses crtc x y
*/
type (
	ModeCursor  C.struct_drm_mode_cursor
	ModeCrtcLut C.struct_drm_mode_crtc_lut
)

const (
	MODE_PAGE_FLIP_EVENT = C.DRM_MODE_PAGE_FLIP_EVENT
	MODE_PAGE_FLIP_FLAGS = C.DRM_MODE_PAGE_FLIP_FLAGS
)

type (
	ModeCrtcPageFlip C.struct_drm_mode_crtc_page_flip

	// create a dumb scanout buffer */
	ModeCreateDumb C.struct_drm_mode_create_dumb

	// set up for mmap of a dumb scanout buffer */
	ModeMapDumb C.struct_drm_mode_map_dumb

	ModeDestroyDumb C.struct_drm_mode_destroy_dumb
)

const (
	IOCTL_BASE                   = C.DRM_IOCTL_BASE
	IOCTL_VERSION                = C.DRM_IOCTL_VERSION
	IOCTL_GET_UNIQUE             = C.DRM_IOCTL_GET_UNIQUE
	IOCTL_GET_MAGIC              = C.DRM_IOCTL_GET_MAGIC
	IOCTL_IRQ_BUSID              = C.DRM_IOCTL_IRQ_BUSID
	IOCTL_GET_MAP                = C.DRM_IOCTL_GET_MAP
	IOCTL_GET_CLIENT             = C.DRM_IOCTL_GET_CLIENT
	IOCTL_GET_STATS              = C.DRM_IOCTL_GET_STATS
	IOCTL_SET_VERSION            = C.DRM_IOCTL_SET_VERSION
	IOCTL_MODESET_CTL            = C.DRM_IOCTL_MODESET_CTL
	IOCTL_GEM_CLOSE              = C.DRM_IOCTL_GEM_CLOSE
	IOCTL_GEM_FLINK              = C.DRM_IOCTL_GEM_FLINK
	IOCTL_GEM_OPEN               = C.DRM_IOCTL_GEM_OPEN
	IOCTL_GET_CAP                = C.DRM_IOCTL_GET_CAP
	IOCTL_SET_UNIQUE             = C.DRM_IOCTL_SET_UNIQUE
	IOCTL_AUTH_MAGIC             = C.DRM_IOCTL_AUTH_MAGIC
	IOCTL_BLOCK                  = C.DRM_IOCTL_BLOCK
	IOCTL_UNBLOCK                = C.DRM_IOCTL_UNBLOCK
	IOCTL_CONTROL                = C.DRM_IOCTL_CONTROL
	IOCTL_ADD_MAP                = C.DRM_IOCTL_ADD_MAP
	IOCTL_ADD_BUFS               = C.DRM_IOCTL_ADD_BUFS
	IOCTL_MARK_BUFS              = C.DRM_IOCTL_MARK_BUFS
	IOCTL_INFO_BUFS              = C.DRM_IOCTL_INFO_BUFS
	IOCTL_MAP_BUFS               = C.DRM_IOCTL_MAP_BUFS
	IOCTL_FREE_BUFS              = C.DRM_IOCTL_FREE_BUFS
	IOCTL_RM_MAP                 = C.DRM_IOCTL_RM_MAP
	IOCTL_SET_SAREA_CTX          = C.DRM_IOCTL_SET_SAREA_CTX
	IOCTL_GET_SAREA_CTX          = C.DRM_IOCTL_GET_SAREA_CTX
	IOCTL_SET_MASTER             = C.DRM_IOCTL_SET_MASTER
	IOCTL_DROP_MASTER            = C.DRM_IOCTL_DROP_MASTER
	IOCTL_ADD_CTX                = C.DRM_IOCTL_ADD_CTX
	IOCTL_RM_CTX                 = C.DRM_IOCTL_RM_CTX
	IOCTL_MOD_CTX                = C.DRM_IOCTL_MOD_CTX
	IOCTL_GET_CTX                = C.DRM_IOCTL_GET_CTX
	IOCTL_SWITCH_CTX             = C.DRM_IOCTL_SWITCH_CTX
	IOCTL_NEW_CTX                = C.DRM_IOCTL_NEW_CTX
	IOCTL_RES_CTX                = C.DRM_IOCTL_RES_CTX
	IOCTL_ADD_DRAW               = C.DRM_IOCTL_ADD_DRAW
	IOCTL_RM_DRAW                = C.DRM_IOCTL_RM_DRAW
	IOCTL_DMA                    = C.DRM_IOCTL_DMA
	IOCTL_LOCK                   = C.DRM_IOCTL_LOCK
	IOCTL_UNLOCK                 = C.DRM_IOCTL_UNLOCK
	IOCTL_FINISH                 = C.DRM_IOCTL_FINISH
	IOCTL_PRIME_HANDLE_TO_FD     = C.DRM_IOCTL_PRIME_HANDLE_TO_FD
	IOCTL_PRIME_FD_TO_HANDLE     = C.DRM_IOCTL_PRIME_FD_TO_HANDLE
	IOCTL_AGP_AQUIRE             = C.DRM_IOCTL_AGP_ACQUIRE
	IOCTL_AGP_RELEASE            = C.DRM_IOCTL_AGP_RELEASE
	IOCTL_AGP_ENABLE             = C.DRM_IOCTL_AGP_ENABLE
	IOCTL_AGP_INFO               = C.DRM_IOCTL_AGP_INFO
	IOCTL_AGP_ALLOC              = C.DRM_IOCTL_AGP_ALLOC
	IOCTL_AGP_FREE               = C.DRM_IOCTL_AGP_FREE
	IOCTL_AGP_BIND               = C.DRM_IOCTL_AGP_BIND
	IOCTL_AGP_UNBIND             = C.DRM_IOCTL_AGP_UNBIND
	IOCTL_SG_ALLOC               = C.DRM_IOCTL_SG_ALLOC
	IOCTL_SG_FREE                = C.DRM_IOCTL_SG_FREE
	IOCTL_WAIT_VBLANK            = C.DRM_IOCTL_WAIT_VBLANK
	IOCTL_UPDATE_DRAW            = C.DRM_IOCTL_UPDATE_DRAW
	IOCTL_MODE_GETRESOURCES      = C.DRM_IOCTL_MODE_GETRESOURCES
	IOCTL_MODE_GETCRTC           = C.DRM_IOCTL_MODE_GETCRTC
	IOCTL_MODE_SETCRTC           = C.DRM_IOCTL_MODE_SETCRTC
	IOCTL_MODE_CURSOR            = C.DRM_IOCTL_MODE_CURSOR
	IOCTL_MODE_GETGAMMA          = C.DRM_IOCTL_MODE_GETGAMMA
	IOCTL_MODE_SETGAMMA          = C.DRM_IOCTL_MODE_SETGAMMA
	IOCTL_MODE_GETENCODER        = C.DRM_IOCTL_MODE_GETENCODER
	IOCTL_MODE_GETCONNECTOR      = C.DRM_IOCTL_MODE_GETCONNECTOR
	IOCTL_MODE_ATTACHMODE        = C.DRM_IOCTL_MODE_ATTACHMODE
	IOCTL_MODE_DETACHMODE        = C.DRM_IOCTL_MODE_DETACHMODE
	IOCTL_MODE_GETPROPERTY       = C.DRM_IOCTL_MODE_GETPROPERTY
	IOCTL_MODE_SETPROPERTY       = C.DRM_IOCTL_MODE_SETPROPERTY
	IOCTL_MODE_GETPROPBLOB       = C.DRM_IOCTL_MODE_GETPROPBLOB
	IOCTL_MODE_GETFB             = C.DRM_IOCTL_MODE_GETFB
	IOCTL_MODE_ADDFB             = C.DRM_IOCTL_MODE_ADDFB
	IOCTL_MODE_RMFB              = C.DRM_IOCTL_MODE_RMFB
	IOCTL_MODE_PAGE_FLIP         = C.DRM_IOCTL_MODE_PAGE_FLIP
	IOCTL_MODE_DIRTYFB           = C.DRM_IOCTL_MODE_DIRTYFB
	IOCTL_MODE_CREATE_DUMB       = C.DRM_IOCTL_MODE_CREATE_DUMB
	IOCTL_MODE_MAP_DUMB          = C.DRM_IOCTL_MODE_MAP_DUMB
	IOCTL_MODE_DESTROY_DUM       = C.DRM_IOCTL_MODE_DESTROY_DUMB
	IOCTL_MODE_GETPLANERESOURCES = C.DRM_IOCTL_MODE_GETPLANERESOURCES
	IOCTL_MODE_GETPLANE          = C.DRM_IOCTL_MODE_GETPLANE
	IOCTL_MODE_SETPLANE          = C.DRM_IOCTL_MODE_SETPLANE
	IOCTL_MODE_ADDFB2            = C.DRM_IOCTL_MODE_ADDFB2
	IOCTL_MODE_OBJ_GETPROPERTIES = C.DRM_IOCTL_MODE_OBJ_GETPROPERTIES
	IOCTL_MODE_OBJ_SETPROPERTY   = C.DRM_IOCTL_MODE_OBJ_SETPROPERTY

	COMMAND_BASE = C.DRM_COMMAND_BASE
	COMMAND_END  = C.DRM_COMMAND_END

	EVENT_VBLANK             = C.DRM_EVENT_VBLANK
	EVENT_FLIP_COMPLETE      = C.DRM_EVENT_FLIP_COMPLETE
	CAP_DUMB_BUFFER          = C.DRM_CAP_DUMB_BUFFER
	CAP_VBLANK_HIGH_CRTC     = C.DRM_CAP_VBLANK_HIGH_CRTC
	CAP_DUMB_PREFERRED_DEPTH = C.DRM_CAP_DUMB_PREFERRED_DEPTH
	CAP_DUMB_PREFER_SHADOW   = C.DRM_CAP_DUMB_PREFER_SHADOW
	CAP_PRIME                = C.DRM_CAP_PRIME
	PRIME_CAP_IMPORT         = C.DRM_PRIME_CAP_IMPORT
	PRIME_CAP_EXPORT         = C.DRM_PRIME_CAP_EXPORT
)
