#include "textflag.h"

// binary_search_arm64.s
TEXT ·BinarySearch(SB), NOSPLIT, $0
    MOVD array+0(FP), R0   // Load array pointer into R0
    MOVD length+8(FP), R1  // Load length into R1
    MOVD target+16(FP), R2 // Load target value into R2
    MOVD $0, R3            // Initialize low index to 0
    MOVD $1, R7
    SUB R1, R7, R4         // Initialize high index to length - 1 (R1 - 1)

binary_search_loop:
    CMP R3, R4             // Compare low and high indices
    BGT binary_search_done // If low > high, target is not present

    // Calculate mid = low + (high - low) / 2
    SUB R4, R3, R5         // high - low in R5
    MOVD $1, R7
    LSR R5, R5, R7         // (high - low) / 2 in R5
    ADD R3, R5, R5         // low + (high - low) / 2 in R5

    // Compare target with array[mid]
    MOVD (R0)(R5<<3), R6   // Load array[mid] into R6
    CMP R2, R6             // Compare target and array[mid]
    BEQ binary_search_found// If equal, target is found
    BLT less_than_mid      // If target is less, search in left half
    BGT greater_than_mid   // If target is greater, search in right half

less_than_mid:
    MOVD $1, R7
    SUB R5, R7, R4         // Set high = mid - 1
    B binary_search_loop

greater_than_mid:
    MOVD $1, R7
    ADD R5, R7, R3         // Set low = mid + 1
    B binary_search_loop

binary_search_found:
    MOVD R5, ret+24(FP)    // Return mid index
    RET

binary_search_done:
    MOVD $-1, R7           // Move -1 to R7
    MOVD R7, ret+24(FP)    // Return -1 (not found)
    RET
