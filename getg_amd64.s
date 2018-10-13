#include "go_asm.h"
#include "textflag.h"

TEXT ·getg(SB), NOSPLIT, $0-32
    MOVQ    TLS, CX
    MOVQ    0(TLS), BX
    MOVQ    (CX)(TLS), DX
    MOVQ    0(CX)(TLS*1), AX
    MOVQ    AX, ret+0(FP)
    MOVQ    BX, ret+8(FP)
    MOVQ    CX, ret+16(FP)
    MOVQ    DX, ret+24(FP)
    RET

TEXT ·getg1(SB), NOSPLIT, $0-8
    MOVQ    FS, CX
    MOVQ    CX, ret+0(FP)
    RET
