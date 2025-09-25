# cudd_Large_4NOV2023.jl
# 24 Oct 2022 @ 2140 hrs
# 13 Nov 2022 @ 2147 hrs modified
# 20 Nov 2023 @ 1935 hrs have begun modifying ps2ns()
# 24 Nov 2023 @ 1542 hrs continue modifying ps2ns()
# 29 Nov 2023 @ 1610 hrs formatting changes to peek()
# 13 Jan 2025 @ 1310 hrs removing fault_B and adding fault to ps2ns()
# 6 Jul 2025 @ 1305 hrs eliminated fault_B.
# 
# pwd    display current directory



# IMPORTS ==============================================================
# ======================================================================

using CUDD
using BenchmarkTools


# BASE SETUP ===========================================================
# ======================================================================

mgr = initialize_cudd()

global ps16 = Cudd_bddNewVar(mgr)
global ps8 = Cudd_bddNewVar(mgr)
global ps4 = Cudd_bddNewVar(mgr)
global ps2 = Cudd_bddNewVar(mgr)
global ps1 = Cudd_bddNewVar(mgr)
global in4 = Cudd_bddNewVar(mgr)
global in2 = Cudd_bddNewVar(mgr)
global in1 = Cudd_bddNewVar(mgr)


function not(a::Ptr{Nothing})
    return Cudd_bddNand(mgr, a, a)
end


function and(a::Ptr{Nothing}, b::Ptr{Nothing})
    return Cudd_bddAnd(mgr, a, b)
end


function or(a::Ptr{Nothing}, b::Ptr{Nothing})
    return Cudd_bddOr(mgr, a, b)
end


function or4(a::Ptr{Nothing}, b::Ptr{Nothing}, c::Ptr{Nothing}, d::Ptr{Nothing})
    return or(a, or(b, or(c, d)))
end

    
# COMMON SETUP =========================================================
# ======================================================================

global nps16 = not(ps16)
global nps8  = not(ps8)
global nps4  = not(ps4)
global nps2  = not(ps2)
global nps1  = not(ps1)

global nin4  = not(in4)
global nin2  = not(in2)
global nin1  = not(in1)

ps16_1 = ps16
ps8_1 = ps8
ps4_1 = ps4
ps2_1 = ps2
ps1_1 = ps1

in4_1 = in4
in2_1 = in2
in1_1 = in1

global all_terms = or(ps1, nps1)
global ONE = all_terms
global no_terms = and(ps1, nps1)
global ZERO = no_terms
global null = no_terms

ps16_s = no_terms
ps8_s = no_terms
ps4_s = no_terms
ps2_s = no_terms
ps1_s = no_terms

in4_s = no_terms
in2_s = no_terms
in1_s = no_terms

global s0  = and(nps16, and(nps8, and(nps4, and(nps2, nps1))))
global s1  = and(nps16, and(nps8, and(nps4, and(nps2, ps1))))
global s2  = and(nps16, and(nps8, and(nps4, and(ps2, nps1))))
global s3  = and(nps16, and(nps8, and(nps4, and(ps2, ps1))))
global s4  = and(nps16, and(nps8, and(ps4, and(nps2, nps1))))
global s5  = and(nps16, and(nps8, and(ps4, and(nps2, ps1))))
global s6  = and(nps16, and(nps8, and(ps4, and(ps2, nps1))))
global s7  = and(nps16, and(nps8, and(ps4, and(ps2, ps1))))
global s8  = and(nps16, and(ps8, and(nps4, and(nps2, nps1))))
global s9  = and(nps16, and(ps8, and(nps4, and(nps2, ps1))))
global s10 = and(nps16, and(ps8, and(nps4, and(ps2, nps1))))
global s11 = and(nps16, and(ps8, and(nps4, and(ps2, ps1))))
global s12 = and(nps16, and(ps8, and(ps4, and(nps2, nps1))))
global s13 = and(nps16, and(ps8, and(ps4, and(nps2, ps1))))
global s14 = and(nps16, and(ps8, and(ps4, and(ps2, nps1))))
global s15 = and(nps16, and(ps8, and(ps4, and(ps2, ps1))))
global s16 = and(ps16, and(nps8, and(nps4, and(nps2, nps1))))
global s17 = and(ps16, and(nps8, and(nps4, and(nps2, ps1))))
global s18 = and(ps16, and(nps8, and(nps4, and(ps2, nps1))))
global s19 = and(ps16, and(nps8, and(nps4, and(ps2, ps1))))
global s20 = and(ps16, and(nps8, and(ps4, and(nps2, nps1))))
global s21 = and(ps16, and(nps8, and(ps4, and(nps2, ps1))))
global s22 = and(ps16, and(nps8, and(ps4, and(ps2, nps1))))
global s23 = and(ps16, and(nps8, and(ps4, and(ps2, ps1))))
global s24 = and(ps16, and(ps8, and(nps4, and(nps2, nps1))))
global s25 = and(ps16, and(ps8, and(nps4, and(nps2, ps1))))
global s26 = and(ps16, and(ps8, and(nps4, and(ps2, nps1))))
global s27 = and(ps16, and(ps8, and(nps4, and(ps2, ps1))))
global s28 = and(ps16, and(ps8, and(ps4, and(nps2, nps1))))
global s29 = and(ps16, and(ps8, and(ps4, and(nps2, ps1))))
global s30 = and(ps16, and(ps8, and(ps4, and(ps2, nps1))))
global s31 = and(ps16, and(ps8, and(ps4, and(ps2, ps1))))

global i0 = and(nin4, and(nin2, nin1))
global i1 = and(nin4, and(nin2, in1))
global i2 = and(nin4, and(in2, nin1))
global i3 = and(nin4, and(in2, in1))
global i4 = and(in4, and(nin2, nin1))
global i5 = and(in4, and(nin2, in1))
global i6 = and(in4, and(in2, nin1))
global i7 = and(in4, and(in2, in1))

global ns0 = not(s0)
global ns1 = not(s1)
global ns2 = not(s2)
global ns3 = not(s3)
global ns4 = not(s4)
global ns5 = not(s5)
global ns6 = not(s6)
global ns7 = not(s7)
global ns8 = not(s8)
global ns9 = not(s9)
global ns10 = not(s10)
global ns11 = not(s11)
global ns12 = not(s12)
global ns13 = not(s13)
global ns14 = not(s14)
global ns15 = not(s15)
global ns16 = not(s16)
global ns17 = not(s17)
global ns18 = not(s18)
global ns19 = not(s19)
global ns20 = not(s20)
global ns21 = not(s21)
global ns22 = not(s22)
global ns23 = not(s23)
global ns24 = not(s24)
global ns25 = not(s25)
global ns26 = not(s26)
global ns27 = not(s27)
global ns28 = not(s28)
global ns29 = not(s29)
global ns30 = not(s30)
global ns31 = not(s31)

global ni0 = not(i0)
global ni1 = not(i1)
global ni2 = not(i2)
global ni3 = not(i3)
global ni4 = not(i4)
global ni5 = not(i5)
global ni6 = not(i6)
global ni7 = not(i7)

fault = "c6:0"
fault_B = String[] # remove later

# tuplejoin ============================================================
# ======================================================================

tuplejoin(t1::Tuple, t2::Tuple, t3...) = tuplejoin((t1..., t2...), t3...)
tuplejoin(t::Tuple) = t


# COMMON LOGIC GATE RULES ==============================================
# ======================================================================

function op1not(s1::Ptr{Nothing}, i1::Ptr{Nothing})
    # computes the propagation function
    o_s::Ptr{Nothing} = s1
    # computes the 1-set
    o_1::Ptr{Nothing} = not(i1)
    return (o_s, o_1)
end


function op2and(s1::Ptr{Nothing}, s2::Ptr{Nothing},
                i1::Ptr{Nothing}, i2::Ptr{Nothing})
    # computes the propagation function
    o_s::Ptr{Nothing} = or4(and(not(i1), and(s1, and(not(i2), s2))),
                            and(i2, and(s1, not(s2))),
                            and(i1, and(s2, not(s1))),
                            and(i1, and(i2, or(s1, s2))))
    # computes the 1-set
    o_1::Ptr{Nothing} = and(i1, i2)
    return (o_s, o_1)
end


function op2or(s1::Ptr{Nothing}, s2::Ptr{Nothing},
               i1::Ptr{Nothing}, i2::Ptr{Nothing})
    # computes the propagation function
    o_s::Ptr{Nothing} = or4(and(i1, and(s1, and(i2, s2))),
                            and(not(i2), and(s1, not(s2))),
                            and(not(i1), and(s2, not(s1))),
                            and(not(i1), and(not(i2), or(s1, s2))))
    # computes the 1-set
    o_1::Ptr{Nothing} = or(i1, i2)
    return (o_s, o_1)
end


function op3and(s1::Ptr{Nothing}, s2::Ptr{Nothing}, s3::Ptr{Nothing},
                i1::Ptr{Nothing}, i2::Ptr{Nothing}, i3::Ptr{Nothing})
    s::Ptr{Nothing}, i::Ptr{Nothing} = op2and(s1, s2, i1, i2)
    o_s::Ptr{Nothing}, o_1::Ptr{Nothing} = op2and(s, s3, i, i3)
    return (o_s, o_1)
end


function op3or(s1::Ptr{Nothing}, s2::Ptr{Nothing}, s3::Ptr{Nothing},
               i1::Ptr{Nothing}, i2::Ptr{Nothing}, i3::Ptr{Nothing})
    s::Ptr{Nothing}, i::Ptr{Nothing} = op2or(s1, s2, i1, i2)
    o_s::Ptr{Nothing}, o_1::Ptr{Nothing} = op2or(s, s3, i, i3)
    return (o_s, o_1)
end


# allSAT ===============================================================
# ======================================================================

function allSAT(f::Ptr{Nothing})::Array{String,1}

    g::Array{String,1} = String[]
   
    # single line conditional statement format
    and(f, and(s0, i0)) != null && push!(g, "s0i0")
    and(f, and(s0, i1)) != null && push!(g, "s0i1")
    and(f, and(s0, i2)) != null && push!(g, "s0i2")
    and(f, and(s0, i3)) != null && push!(g, "s0i3")
    and(f, and(s0, i4)) != null && push!(g, "s0i4")
    and(f, and(s0, i5)) != null && push!(g, "s0i5")
    and(f, and(s0, i6)) != null && push!(g, "s0i6")
    and(f, and(s0, i7)) != null && push!(g, "s0i7")
    and(f, and(s1, i0)) != null && push!(g, "s1i0")
    and(f, and(s1, i1)) != null && push!(g, "s1i1")
    and(f, and(s1, i2)) != null && push!(g, "s1i2")
    and(f, and(s1, i3)) != null && push!(g, "s1i3")
    and(f, and(s1, i4)) != null && push!(g, "s1i4")
    and(f, and(s1, i5)) != null && push!(g, "s1i5")
    and(f, and(s1, i6)) != null && push!(g, "s1i6")
    and(f, and(s1, i7)) != null && push!(g, "s1i7")
    and(f, and(s2, i0)) != null && push!(g, "s2i0")
    and(f, and(s2, i1)) != null && push!(g, "s2i1")
    and(f, and(s2, i2)) != null && push!(g, "s2i2")
    and(f, and(s2, i3)) != null && push!(g, "s2i3")
    and(f, and(s2, i4)) != null && push!(g, "s2i4")
    and(f, and(s2, i5)) != null && push!(g, "s2i5")
    and(f, and(s2, i6)) != null && push!(g, "s2i6")
    and(f, and(s2, i7)) != null && push!(g, "s2i7")
    and(f, and(s3, i0)) != null && push!(g, "s3i0")
    and(f, and(s3, i1)) != null && push!(g, "s3i1")
    and(f, and(s3, i2)) != null && push!(g, "s3i2")
    and(f, and(s3, i3)) != null && push!(g, "s3i3")
    and(f, and(s3, i4)) != null && push!(g, "s3i4")
    and(f, and(s3, i5)) != null && push!(g, "s3i5")
    and(f, and(s3, i6)) != null && push!(g, "s3i6")
    and(f, and(s3, i7)) != null && push!(g, "s3i7")
    and(f, and(s4, i0)) != null && push!(g, "s4i0")
    and(f, and(s4, i1)) != null && push!(g, "s4i1")
    and(f, and(s4, i2)) != null && push!(g, "s4i2")
    and(f, and(s4, i3)) != null && push!(g, "s4i3")
    and(f, and(s4, i4)) != null && push!(g, "s4i4")
    and(f, and(s4, i5)) != null && push!(g, "s4i5")
    and(f, and(s4, i6)) != null && push!(g, "s4i6")
    and(f, and(s4, i7)) != null && push!(g, "s4i7")
    and(f, and(s5, i0)) != null && push!(g, "s5i0")
    and(f, and(s5, i1)) != null && push!(g, "s5i1")
    and(f, and(s5, i2)) != null && push!(g, "s5i2")
    and(f, and(s5, i3)) != null && push!(g, "s5i3")
    and(f, and(s5, i4)) != null && push!(g, "s5i4")
    and(f, and(s5, i5)) != null && push!(g, "s5i5")
    and(f, and(s5, i6)) != null && push!(g, "s5i6")
    and(f, and(s5, i7)) != null && push!(g, "s5i7")
    and(f, and(s6, i0)) != null && push!(g, "s6i0")
    and(f, and(s6, i1)) != null && push!(g, "s6i1")
    and(f, and(s6, i2)) != null && push!(g, "s6i2")
    and(f, and(s6, i3)) != null && push!(g, "s6i3")
    and(f, and(s6, i4)) != null && push!(g, "s6i4")
    and(f, and(s6, i5)) != null && push!(g, "s6i5")
    and(f, and(s6, i6)) != null && push!(g, "s6i6")
    and(f, and(s6, i7)) != null && push!(g, "s6i7")
    and(f, and(s7, i0)) != null && push!(g, "s7i0")
    and(f, and(s7, i1)) != null && push!(g, "s7i1")
    and(f, and(s7, i2)) != null && push!(g, "s7i2")
    and(f, and(s7, i3)) != null && push!(g, "s7i3")
    and(f, and(s7, i4)) != null && push!(g, "s7i4")
    and(f, and(s7, i5)) != null && push!(g, "s7i5")
    and(f, and(s7, i6)) != null && push!(g, "s7i6")
    and(f, and(s7, i7)) != null && push!(g, "s8i7")
    and(f, and(s8, i0)) != null && push!(g, "s8i0")
    and(f, and(s8, i1)) != null && push!(g, "s8i1")
    and(f, and(s8, i2)) != null && push!(g, "s8i2")
    and(f, and(s8, i3)) != null && push!(g, "s8i3")
    and(f, and(s8, i4)) != null && push!(g, "s8i4")
    and(f, and(s8, i5)) != null && push!(g, "s8i5")
    and(f, and(s8, i6)) != null && push!(g, "s8i6")
    and(f, and(s8, i7)) != null && push!(g, "s8i7")
    and(f, and(s9, i0)) != null && push!(g, "s9i0")
    and(f, and(s9, i1)) != null && push!(g, "s9i1")
    and(f, and(s9, i2)) != null && push!(g, "s9i2")
    and(f, and(s9, i3)) != null && push!(g, "s9i3")
    and(f, and(s9, i4)) != null && push!(g, "s9i4")
    and(f, and(s9, i5)) != null && push!(g, "s9i5")
    and(f, and(s9, i6)) != null && push!(g, "s9i6")
    and(f, and(s9, i7)) != null && push!(g, "s9i7")
    and(f, and(s10, i0)) != null && push!(g, "s10i0")
    and(f, and(s10, i1)) != null && push!(g, "s10i1")
    and(f, and(s10, i2)) != null && push!(g, "s10i2")
    and(f, and(s10, i3)) != null && push!(g, "s10i3")
    and(f, and(s10, i4)) != null && push!(g, "s10i4")
    and(f, and(s10, i5)) != null && push!(g, "s10i5")
    and(f, and(s10, i6)) != null && push!(g, "s10i6")
    and(f, and(s10, i7)) != null && push!(g, "s10i7")
    and(f, and(s11, i0)) != null && push!(g, "s11i0")
    and(f, and(s11, i1)) != null && push!(g, "s11i1")
    and(f, and(s11, i2)) != null && push!(g, "s11i2")
    and(f, and(s11, i3)) != null && push!(g, "s11i3")
    and(f, and(s11, i4)) != null && push!(g, "s11i4")
    and(f, and(s11, i5)) != null && push!(g, "s11i5")
    and(f, and(s11, i6)) != null && push!(g, "s11i6")
    and(f, and(s11, i7)) != null && push!(g, "s11i7")
    and(f, and(s12, i0)) != null && push!(g, "s12i0")
    and(f, and(s12, i1)) != null && push!(g, "s12i1")
    and(f, and(s12, i2)) != null && push!(g, "s12i2")
    and(f, and(s12, i3)) != null && push!(g, "s12i3")
    and(f, and(s12, i4)) != null && push!(g, "s12i4")
    and(f, and(s12, i5)) != null && push!(g, "s12i5")
    and(f, and(s12, i6)) != null && push!(g, "s12i6")
    and(f, and(s12, i7)) != null && push!(g, "s12i7")
    and(f, and(s13, i0)) != null && push!(g, "s13i0")
    and(f, and(s13, i1)) != null && push!(g, "s13i1")
    and(f, and(s13, i2)) != null && push!(g, "s13i2")
    and(f, and(s13, i3)) != null && push!(g, "s13i3")
    and(f, and(s13, i4)) != null && push!(g, "s13i4")
    and(f, and(s13, i5)) != null && push!(g, "s13i5")
    and(f, and(s13, i6)) != null && push!(g, "s13i6")
    and(f, and(s13, i7)) != null && push!(g, "s13i7")
    and(f, and(s14, i0)) != null && push!(g, "s14i0")
    and(f, and(s14, i1)) != null && push!(g, "s14i1")
    and(f, and(s14, i2)) != null && push!(g, "s14i2")
    and(f, and(s14, i3)) != null && push!(g, "s14i3")
    and(f, and(s14, i4)) != null && push!(g, "s14i4")
    and(f, and(s14, i5)) != null && push!(g, "s14i5")
    and(f, and(s14, i6)) != null && push!(g, "s14i6")
    and(f, and(s14, i7)) != null && push!(g, "s14i7")
    and(f, and(s15, i0)) != null && push!(g, "s15i0")
    and(f, and(s15, i1)) != null && push!(g, "s15i1")
    and(f, and(s15, i2)) != null && push!(g, "s15i2")
    and(f, and(s15, i3)) != null && push!(g, "s15i3")
    and(f, and(s15, i4)) != null && push!(g, "s15i4")
    and(f, and(s15, i5)) != null && push!(g, "s15i5")
    and(f, and(s15, i6)) != null && push!(g, "s15i6")
    and(f, and(s15, i7)) != null && push!(g, "s15i7")
    and(f, and(s16, i0)) != null && push!(g, "s16i0")
    and(f, and(s16, i1)) != null && push!(g, "s16i1")
    and(f, and(s16, i2)) != null && push!(g, "s16i2")
    and(f, and(s16, i3)) != null && push!(g, "s16i3")
    and(f, and(s16, i4)) != null && push!(g, "s16i4")
    and(f, and(s16, i5)) != null && push!(g, "s16i5")
    and(f, and(s16, i6)) != null && push!(g, "s16i6")
    and(f, and(s16, i7)) != null && push!(g, "s16i7")
    and(f, and(s17, i0)) != null && push!(g, "s17i0")
    and(f, and(s17, i1)) != null && push!(g, "s17i1")
    and(f, and(s17, i2)) != null && push!(g, "s17i2")
    and(f, and(s17, i3)) != null && push!(g, "s17i3")
    and(f, and(s17, i4)) != null && push!(g, "s17i4")
    and(f, and(s17, i5)) != null && push!(g, "s17i5")
    and(f, and(s17, i6)) != null && push!(g, "s17i6")
    and(f, and(s17, i7)) != null && push!(g, "s17i7")
    and(f, and(s18, i0)) != null && push!(g, "s18i0")
    and(f, and(s18, i1)) != null && push!(g, "s18i1")
    and(f, and(s18, i2)) != null && push!(g, "s18i2")
    and(f, and(s18, i3)) != null && push!(g, "s18i3")
    and(f, and(s18, i4)) != null && push!(g, "s18i4")
    and(f, and(s18, i5)) != null && push!(g, "s18i5")
    and(f, and(s18, i6)) != null && push!(g, "s18i6")
    and(f, and(s18, i7)) != null && push!(g, "s18i7")
    and(f, and(s19, i0)) != null && push!(g, "s19i0")
    and(f, and(s19, i1)) != null && push!(g, "s19i1")
    and(f, and(s19, i2)) != null && push!(g, "s19i2")
    and(f, and(s19, i3)) != null && push!(g, "s19i3")
    and(f, and(s19, i4)) != null && push!(g, "s19i4")
    and(f, and(s19, i5)) != null && push!(g, "s19i5")
    and(f, and(s19, i6)) != null && push!(g, "s19i6")
    and(f, and(s19, i7)) != null && push!(g, "s19i7")
    and(f, and(s20, i0)) != null && push!(g, "s20i0")
    and(f, and(s20, i1)) != null && push!(g, "s20i1")
    and(f, and(s20, i2)) != null && push!(g, "s20i2")
    and(f, and(s20, i3)) != null && push!(g, "s20i3")
    and(f, and(s20, i4)) != null && push!(g, "s20i4")
    and(f, and(s20, i5)) != null && push!(g, "s20i5")
    and(f, and(s20, i6)) != null && push!(g, "s20i6")
    and(f, and(s20, i7)) != null && push!(g, "s20i7")
    and(f, and(s21, i0)) != null && push!(g, "s21i0")
    and(f, and(s21, i1)) != null && push!(g, "s21i1")
    and(f, and(s21, i2)) != null && push!(g, "s21i2")
    and(f, and(s21, i3)) != null && push!(g, "s21i3")
    and(f, and(s21, i4)) != null && push!(g, "s21i4")
    and(f, and(s21, i5)) != null && push!(g, "s21i5")
    and(f, and(s21, i6)) != null && push!(g, "s21i6")
    and(f, and(s21, i7)) != null && push!(g, "s21i7")
    and(f, and(s22, i0)) != null && push!(g, "s22i0")
    and(f, and(s22, i1)) != null && push!(g, "s22i1")
    and(f, and(s22, i2)) != null && push!(g, "s22i2")
    and(f, and(s22, i3)) != null && push!(g, "s22i3")
    and(f, and(s22, i4)) != null && push!(g, "s22i4")
    and(f, and(s22, i5)) != null && push!(g, "s22i5")
    and(f, and(s22, i6)) != null && push!(g, "s22i6")
    and(f, and(s22, i7)) != null && push!(g, "s22i7")
    and(f, and(s23, i0)) != null && push!(g, "s23i0")
    and(f, and(s23, i1)) != null && push!(g, "s23i1")
    and(f, and(s23, i2)) != null && push!(g, "s23i2")
    and(f, and(s23, i3)) != null && push!(g, "s23i3")
    and(f, and(s23, i4)) != null && push!(g, "s23i4")
    and(f, and(s23, i5)) != null && push!(g, "s23i5")
    and(f, and(s23, i6)) != null && push!(g, "s23i6")
    and(f, and(s23, i7)) != null && push!(g, "s23i7")
    and(f, and(s24, i0)) != null && push!(g, "s24i0")
    and(f, and(s24, i1)) != null && push!(g, "s24i1")
    and(f, and(s24, i2)) != null && push!(g, "s24i2")
    and(f, and(s24, i3)) != null && push!(g, "s24i3")
    and(f, and(s24, i4)) != null && push!(g, "s24i4")
    and(f, and(s24, i5)) != null && push!(g, "s24i5")
    and(f, and(s24, i6)) != null && push!(g, "s24i6")
    and(f, and(s24, i7)) != null && push!(g, "s24i7")
    and(f, and(s25, i0)) != null && push!(g, "s25i0")
    and(f, and(s25, i1)) != null && push!(g, "s25i1")
    and(f, and(s25, i2)) != null && push!(g, "s25i2")
    and(f, and(s25, i3)) != null && push!(g, "s25i3")
    and(f, and(s25, i4)) != null && push!(g, "s25i4")
    and(f, and(s25, i5)) != null && push!(g, "s25i5")
    and(f, and(s25, i6)) != null && push!(g, "s25i6")
    and(f, and(s25, i7)) != null && push!(g, "s25i7")
    and(f, and(s26, i0)) != null && push!(g, "s26i0")
    and(f, and(s26, i1)) != null && push!(g, "s26i1")
    and(f, and(s26, i2)) != null && push!(g, "s26i2")
    and(f, and(s26, i3)) != null && push!(g, "s26i3")
    and(f, and(s26, i4)) != null && push!(g, "s26i4")
    and(f, and(s26, i5)) != null && push!(g, "s26i5")
    and(f, and(s26, i6)) != null && push!(g, "s26i6")
    and(f, and(s26, i7)) != null && push!(g, "s26i7")
    and(f, and(s27, i0)) != null && push!(g, "s27i0")
    and(f, and(s27, i1)) != null && push!(g, "s27i1")
    and(f, and(s27, i2)) != null && push!(g, "s27i2")
    and(f, and(s27, i3)) != null && push!(g, "s27i3")
    and(f, and(s27, i4)) != null && push!(g, "s27i4")
    and(f, and(s27, i5)) != null && push!(g, "s27i5")
    and(f, and(s27, i6)) != null && push!(g, "s27i6")
    and(f, and(s27, i7)) != null && push!(g, "s27i7")
    and(f, and(s28, i0)) != null && push!(g, "s28i0")
    and(f, and(s28, i1)) != null && push!(g, "s28i1")
    and(f, and(s28, i2)) != null && push!(g, "s28i2")
    and(f, and(s28, i3)) != null && push!(g, "s28i3")
    and(f, and(s28, i4)) != null && push!(g, "s28i4")
    and(f, and(s28, i5)) != null && push!(g, "s28i5")
    and(f, and(s28, i6)) != null && push!(g, "s28i6")
    and(f, and(s28, i7)) != null && push!(g, "s28i7")
    and(f, and(s29, i0)) != null && push!(g, "s29i0")
    and(f, and(s29, i1)) != null && push!(g, "s29i1")
    and(f, and(s29, i2)) != null && push!(g, "s29i2")
    and(f, and(s29, i3)) != null && push!(g, "s29i3")
    and(f, and(s29, i4)) != null && push!(g, "s29i4")
    and(f, and(s29, i5)) != null && push!(g, "s29i5")
    and(f, and(s29, i6)) != null && push!(g, "s29i6")
    and(f, and(s29, i7)) != null && push!(g, "s29i7")
    and(f, and(s30, i0)) != null && push!(g, "s30i0")
    and(f, and(s30, i1)) != null && push!(g, "s30i1")
    and(f, and(s30, i2)) != null && push!(g, "s30i2")
    and(f, and(s30, i3)) != null && push!(g, "s30i3")
    and(f, and(s30, i4)) != null && push!(g, "s30i4")
    and(f, and(s30, i5)) != null && push!(g, "s30i5")
    and(f, and(s30, i6)) != null && push!(g, "s30i6")
    and(f, and(s30, i7)) != null && push!(g, "s30i7")
    and(f, and(s31, i0)) != null && push!(g, "s31i0")
    and(f, and(s31, i1)) != null && push!(g, "s31i1")
    and(f, and(s31, i2)) != null && push!(g, "s31i2")
    and(f, and(s31, i3)) != null && push!(g, "s31i3")
    and(f, and(s31, i4)) != null && push!(g, "s31i4")
    and(f, and(s31, i5)) != null && push!(g, "s31i5")
    and(f, and(s31, i6)) != null && push!(g, "s31i6")
    and(f, and(s31, i7)) != null && push!(g, "s31i7")
    
    return g

end
        

# ps2ns ================================================================
# ======================================================================

function ps2ns(ps16_s::Ptr{Nothing}, ps8_s::Ptr{Nothing}, ps4_s::Ptr{Nothing},
     ps2_s::Ptr{Nothing}, ps1_s::Ptr{Nothing}, fault::String)

    # level 1

    # -----------------------------------------------------------
    nps1_s, nps1_1 = op1not(ps1_s, ps1_1)
    fault == "nps1:0" && (nps1_s = nps1_1)
    fault == "nps1:1" && (nps1_s = not(nps1_1))
    # -----------------------------------------------------------
    nps2_s, nps2_1 = op1not(ps2_s, ps2_1)
    fault == "nps2:0" && (nps2_s = nps2_1)
    fault == "nps2:1" && (nps2_s = not(nps2_1))
    # -----------------------------------------------------------
    nps4_s, nps4_1 = op1not(ps4_s, ps4_1)
    fault == "nps4:0" && (nps4_s = nps4_1)
    fault == "nps4:1" && (nps4_s = not(nps4_1))
    # -----------------------------------------------------------
    nps8_s, nps8_1 = op1not(ps8_s, ps8_1)
    fault == "nps8:0" && (nps8_s = nps8_1)
    fault == "nps8:1" && (nps8_s = not(nps8_1))
    # -----------------------------------------------------------
    nps16_s, nps16_1 = op1not(ps16_s,ps16_1)
    fault == "nps16:0" && (nps16_s = nps16_1)
    fault == "nps16:1" && (nps16_s = not(nps16_1))
    # -----------------------------------------------------------
    nin1_s, nin1_1 = op1not(in1_s, in1_1)
    fault == "nin1:0" && (nin1_s = nin1_1)
    fault == "nin1:1" && (nin1_s = not(nin1_1))
    # -----------------------------------------------------------
    nin2_s, nin2_1 = op1not(in2_s, in2_1)
    fault == "nin2:0" && (nin2_s = nin2_1)
    fault == "nin2:1" && (nin2_s = not(nin2_1))
    # -----------------------------------------------------------
    nin4_s, nin4_1 = op1not(in4_s, in4_1)
    fault == "nin4:0" && (nin4_s = nin4_1)
    fault == "nin4:1" && (nin4_s = not(nin4_1))
    # -----------------------------------------------------------
    i7_s, i7_1 = op3and(in4_s,in2_s,in1_s,in4_1,in2_1,in1_1)
    fault == "i7:0" && (i7_s = i7_1)
    fault == "i7:1" && (i7_s = not(i7_1))
    # -----------------------------------------------------------
    
    # level 2

    # -----------------------------------------------------------
    i0_s, i0_1 = op3and(nin4_s,nin2_s,nin1_s,nin4_1,nin2_1,nin1_1)
    fault == "i0:0" && (i0_s = i0_1)
    fault == "i0:1" && (i0_s = not(i0_1))
    # -----------------------------------------------------------
    i1_s, i1_1 = op3and(nin4_s,nin2_s,in1_s,nin4_1,nin2_1,in1_1)
    fault == "i1:0" && (i1_s = i1_1)
    fault == "i1:1" && (i1_s = not(i1_1))
    # -----------------------------------------------------------
    i2_s, i2_1 = op3and(nin4_s,in2_s,nin1_s,nin4_1,in2_1,nin1_1)
    fault == "i2:0" && (i2_s = i2_1)
    fault == "i2:1" && (i2_s = not(i2_1))
    # -----------------------------------------------------------
    i3_s, i3_1 = op3and(nin4_s,in2_s,in1_s,nin4_1,in2_1,in1_1)
    fault == "i3:0" && (i3_s = i3_1)
    fault == "i3:1" && (i3_s = not(i3_1))
    # -----------------------------------------------------------
    i4_s, i4_1 = op3and(in4_s,nin2_s,nin1_s,in4_1,nin2_1,nin1_1)
    fault == "i4:0" && (i4_s = i4_1)
    fault == "i4:1" && (i4_s = not(i4_1))
    # -----------------------------------------------------------
    i5_s, i5_1 = op3and(in4_s,nin2_s,in1_s,in4_1,nin2_1,in1_1)
    fault == "i5:0" && (i5_s = i5_1)
    fault == "i5:1" && (i5_s = not(i5_1))
    # -----------------------------------------------------------
    i6_s, i6_1 = op3and(in4_s,in2_s,nin1_s,in4_1,in2_1,nin1_1)
    fault == "i6:0" && (i6_s = i6_1)
    fault == "i6:1" && (i6_s = not(i6_1))
    # -----------------------------------------------------------
    ls0_s, ls0_1 = op3and(nps4_s,nps2_s,nps1_s,nps4_1,nps2_1,nps1_1)
    fault == "ls0:0" && (ls0_s = ls0_1)
    fault == "ls0:1" && (ls0_s = not(ls0_1))
    # -----------------------------------------------------------
    ls1_s, ls1_1 = op3and(nps4_s,nps2_s,ps1_s,nps4_1,nps2_1,ps1_1)
    fault == "ls1:0" && (ls1_s = ls1_1)
    fault == "ls1:1" && (ls1_s = not(ls1_1))
    # -----------------------------------------------------------
    ls2_s, ls2_1 = op3and(nps4_s,ps2_s,nps1_s,nps4_1,ps2_1,nps1_1)
    fault == "ls2:0" && (ls2_s = ls2_1)
    fault == "ls2:1" && (ls2_s = not(ls2_1))
    # -----------------------------------------------------------
    ls3_s, ls3_1 = op3and(nps4_s,ps2_s,ps1_s,nps4_1,ps2_1,ps1_1)
    fault == "ls3:0" && (ls3_s = ls3_1)
    fault == "ls3:1" && (ls3_s = not(ls3_1))
    # -----------------------------------------------------------
    ls4_s, ls4_1 = op3and(ps4_s,nps2_s,nps1_s,ps4_1,nps2_1,nps1_1)
    fault == "ls4:0" && (ls4_s = ls4_1)
    fault == "ls4:1" && (ls4_s = not(ls4_1))
    # -----------------------------------------------------------
    ls5_s, ls5_1 = op3and(ps4_s,nps2_s,ps1_s,ps4_1,nps2_1,ps1_1)
    fault == "ls5:0" && (ls5_s = ls5_1)
    fault == "ls5:1" && (ls5_s = not(ls5_1))
    # -----------------------------------------------------------
    ls6_s, ls6_1 = op3and(ps4_s,ps2_s,nps1_s,ps4_1,ps2_1,nps1_1)
    fault == "ls6:0" && (ls6_s = ls6_1)
    fault == "ls6:1" && (ls6_s = not(ls6_1))
    # -----------------------------------------------------------
    ls7_s, ls7_1 = op3and(ps4_s,ps2_s,ps1_s,ps4_1,ps2_1,ps1_1)
    fault == "ls7:0" && (ls7_s = ls7_1)
    fault == "ls7:1" && (ls7_s = not(ls7_1))
    # -----------------------------------------------------------
    ni7_s, ni7_1 = op1not(i7_s,i7_1)
    fault == "ni7:0" && (ni7_s = ni7_1)
    fault == "ni7:1" && (ni7_s = not(ni7_1))
    # -----------------------------------------------------------
    s31_s, s31_1 = op3and(ps16_s,ps8_s,ls7_s,ps16_1,ps8_1,ls7_1)
    fault == "s31:0" && (s31_s = s31_1)
    fault == "s31:1" && (s31_s = not(s31_1))
    # -----------------------------------------------------------

    # level 3

    # -----------------------------------------------------------
    ni0_s, ni0_1 = op1not(i0_s,i0_1)
    fault == "ni0:0" && (ni0_s = ni0_1)
    fault == "ni0:1" && (ni0_s = not(ni0_1))
    # -----------------------------------------------------------
    ni1_s, ni1_1 = op1not(i1_s,i1_1)
    fault == "ni1:0" && (ni1_s = ni1_1)
    fault == "ni1:1" && (ni1_s = not(ni1_1))
    # -----------------------------------------------------------
    ni2_s, ni2_1 = op1not(i2_s,i2_1)
    fault == "ni2:0" && (ni2_s = ni2_1)
    fault == "ni2:1" && (ni2_s = not(ni2_1))
    # -----------------------------------------------------------
    ni3_s, ni3_1 = op1not(i3_s,i3_1)
    fault == "ni3:0" && (ni3_s = ni3_1)
    fault == "ni3:1" && (ni3_s = not(ni3_1))
    # -----------------------------------------------------------
    ni5_s, ni5_1 = op1not(i5_s,i5_1)
    fault == "ni5:0" && (ni5_s = ni5_1)
    fault == "ni5:1" && (ni5_s = not(ni5_1))
    # -----------------------------------------------------------
    ni6_s, ni6_1 = op1not(i6_s,i6_1)
    fault == "ni6:0" && (ni6_s = ni6_1)
    fault == "ni6:1" && (ni6_s = not(ni6_1))
    # -----------------------------------------------------------
    s0_s, s0_1 = op3and(nps16_s,nps8_s,ls0_s,nps16_1,nps8_1,ls0_1)
    fault == "s0:0" && (s0_s = s0_1)
    fault == "s0:1" && (s0_s = not(s0_1))
    # -----------------------------------------------------------
    s1_s, s1_1 = op3and(nps16_s,nps8_s,ls1_s,nps16_1,nps8_1,ls1_1)
    fault == "s1:0" && (s1_s = s1_1)
    fault == "s1:1" && (s1_s = not(s1_1))
    # -----------------------------------------------------------
    s2_s, s2_1 = op3and(nps16_s,nps8_s,ls2_s,nps16_1,nps8_1,ls2_1)
    fault == "s2:0" && (s2_s = s2_1)
    fault == "s2:1" && (s2_s = not(s2_1))
    # -----------------------------------------------------------
    s3_s, s3_1 = op3and(nps16_s,nps8_s,ls3_s,nps16_1,nps8_1,ls3_1)
    fault == "s3:0" && (s3_s = s3_1)
    fault == "s3:1" && (s3_s = not(s3_1))
    # -----------------------------------------------------------
    s4_s, s4_1 = op3and(nps16_s,nps8_s,ls4_s,nps16_1,nps8_1,ls4_1)
    fault == "s4:0" && (s4_s = s4_1)
    fault == "s4:1" && (s4_s = not(s4_1))
    # -----------------------------------------------------------
    s5_s, s5_1 = op3and(nps16_s,nps8_s,ls5_s,nps16_1,nps8_1,ls5_1)
    fault == "s5:0" && (s5_s = s5_1)
    fault == "s5:1" && (s5_s = not(s5_1))
    # -----------------------------------------------------------
    s6_s, s6_1 = op3and(nps16_s,nps8_s,ls6_s,nps16_1,nps8_1,ls6_1)
    fault == "s6:0" && (s6_s = s6_1)
    fault == "s6:1" && (s6_s = not(s6_1))
    # -----------------------------------------------------------
    s7_s, s7_1 = op3and(nps16_s,nps8_s,ls7_s,nps16_1,nps8_1,ls7_1)
    fault == "s7:0" && (s7_s = s7_1)
    fault == "s7:1" && (s7_s = not(s7_1))
    # -----------------------------------------------------------
    s8_s, s8_1 = op3and(nps16_s,ps8_s,ls0_s,nps16_1,ps8_1,ls0_1)
    fault == "s8:0" && (s8_s = s8_1)
    fault == "s8:1" && (s8_s = not(s8_1))
    # -----------------------------------------------------------
    s9_s, s9_1 = op3and(nps16_s,ps8_s,ls1_s,nps16_1,ps8_1,ls1_1)
    fault == "s9:0" && (s9_s = s9_1)
    fault == "s9:1" && (s9_s = not(s9_1))
    # -----------------------------------------------------------
    s10_s, s10_1 = op3and(nps16_s,ps8_s,ls2_s,nps16_1,ps8_1,ls2_1)
    fault == "s10:0" && (s10_s = s10_1)
    fault == "s10:1" && (s10_s = not(s10_1))
    # -----------------------------------------------------------
    s11_s, s11_1 = op3and(nps16_s,ps8_s,ls3_s,nps16_1,ps8_1,ls3_1)
    fault == "s11:0" && (s11_s = s11_1)
    fault == "s11:1" && (s11_s = not(s11_1))
    # -----------------------------------------------------------
    s12_s, s12_1 = op3and(nps16_s,ps8_s,ls4_s,nps16_1,ps8_1,ls4_1)
    fault == "s12:0" && (s12_s = s12_1)
    fault == "s12:1" && (s12_s = not(s12_1))
    # -----------------------------------------------------------
    s13_s, s13_1 = op3and(nps16_s,ps8_s,ls5_s,nps16_1,ps8_1,ls5_1)
    fault == "s13:0" && (s13_s = s13_1)
    fault == "s13:1" && (s13_s = not(s13_1))
    # -----------------------------------------------------------
    s14_s, s14_1 = op3and(nps16_s,ps8_s,ls6_s,nps16_1,ps8_1,ls6_1)
    fault == "s14:0" && (s14_s = s14_1)
    fault == "s14:1" && (s14_s = not(s14_1))
    # -----------------------------------------------------------
    s15_s, s15_1 = op3and(nps16_s,ps8_s,ls7_s,nps16_1,ps8_1,ls7_1)
    fault == "s15:0" && (s15_s = s15_1)
    fault == "s15:1" && (s15_s = not(s15_1))
    # -----------------------------------------------------------
    s16_s, s16_1 = op3and(ps16_s,nps8_s,ls0_s,ps16_1,nps8_1,ls0_1)
    fault == "s16:0" && (s16_s = s16_1)
    fault == "s16:1" && (s16_s = not(s16_1))
    # -----------------------------------------------------------
    s17_s, s17_1 = op3and(ps16_s,nps8_s,ls1_s,ps16_1,nps8_1,ls1_1)
    fault == "s17:0" && (s17_s = s17_1)
    fault == "s17:1" && (s17_s = not(s17_1))
    # -----------------------------------------------------------
    s18_s, s18_1 = op3and(ps16_s,nps8_s,ls2_s,ps16_1,nps8_1,ls2_1)
    fault == "s18:0" && (s18_s = s18_1)
    fault == "s18:1" && (s18_s = not(s18_1))
    # -----------------------------------------------------------
    s19_s, s19_1 = op3and(ps16_s,nps8_s,ls3_s,ps16_1,nps8_1,ls3_1)
    fault == "s19:0" && (s19_s = s19_1)
    fault == "s19:1" && (s19_s = not(s19_1))
    # -----------------------------------------------------------
    s20_s, s20_1 = op3and(ps16_s,nps8_s,ls4_s,ps16_1,nps8_1,ls4_1)
    fault == "s20:0" && (s20_s = s20_1)
    fault == "s20:1" && (s20_s = not(s20_1))
    # -----------------------------------------------------------
    s21_s, s21_1 = op3and(ps16_s,nps8_s,ls5_s,ps16_1,nps8_1,ls5_1)
    fault == "s21:0" && (s21_s = s21_1)
    fault == "s21:1" && (s21_s = not(s21_1))
    # -----------------------------------------------------------
    s22_s, s22_1 = op3and(ps16_s,nps8_s,ls6_s,ps16_1,nps8_1,ls6_1)
    fault == "s22:0" && (s22_s = s22_1)
    fault == "s22:1" && (s22_s = not(s22_1))
    # -----------------------------------------------------------
    s23_s, s23_1 = op3and(ps16_s,nps8_s,ls7_s,ps16_1,nps8_1,ls7_1)
    fault == "s23:0" && (s23_s = s23_1)
    fault == "s23:1" && (s23_s = not(s23_1))
    # -----------------------------------------------------------
    s24_s, s24_1 = op3and(ps16_s,ps8_s,ls0_s,ps16_1,ps8_1,ls0_1)
    fault == "s24:0" && (s24_s = s24_1)
    fault == "s24:1" && (s24_s = not(s24_1))
    # -----------------------------------------------------------
    s25_s, s25_1 = op3and(ps16_s,ps8_s,ls1_s,ps16_1,ps8_1,ls1_1)
    fault == "s25:0" && (s25_s = s25_1)
    fault == "s25:1" && (s25_s = not(s25_1))
    # -----------------------------------------------------------
    s26_s, s26_1 = op3and(ps16_s,ps8_s,ls2_s,ps16_1,ps8_1,ls2_1)
    fault == "s26:0" && (s26_s = s26_1)
    fault == "s26:1" && (s26_s = not(s26_1))
    # -----------------------------------------------------------
    s27_s, s27_1 = op3and(ps16_s,ps8_s,ls3_s,ps16_1,ps8_1,ls3_1)
    fault == "s27:0" && (s27_s = s27_1)
    fault == "s27:1" && (s27_s = not(s27_1))
    # -----------------------------------------------------------
    s28_s, s28_1 = op3and(ps16_s,ps8_s,ls4_s,ps16_1,ps8_1,ls4_1)
    fault == "s28:0" && (s28_s = s28_1)
    fault == "s28:1" && (s28_s = not(s28_1))
    # -----------------------------------------------------------
    s29_s, s29_1 = op3and(ps16_s,ps8_s,ls5_s,ps16_1,ps8_1,ls5_1)
    fault == "s29:0" && (s29_s = s29_1)
    fault == "s29:1" && (s29_s = not(s29_1))
    # -----------------------------------------------------------
    s30_s, s30_1 = op3and(ps16_s,ps8_s,ls6_s,ps16_1,ps8_1,ls6_1)
    fault == "s30:0" && (s30_s = s30_1)
    fault == "s30:1" && (s30_s = not(s30_1))
    # -----------------------------------------------------------
    b2_s, b2_1 = op3or(i5_s,i3_s,i2_s,i5_1,i3_1,i2_1)
    fault == "b2:0" && (b2_s = b2_1)
    fault == "b2:1" && (b2_s = not(b2_1))
    # -----------------------------------------------------------
    b7_s, b7_1 = op2or(i5_s,i1_s,i5_1,i1_1)
    fault == "b7:0" && (b7_s = b7_1)
    fault == "b7:1" && (b7_s = not(b7_1))
    # -----------------------------------------------------------
    c5_s, c5_1 = op2or(i7_s,i5_s,i7_1,i5_1)
    fault == "c5:0" && (c5_s = c5_1)
    fault == "c5:1" && (c5_s = not(c5_1))
    # -----------------------------------------------------------
    c13_s, c13_1 = op2or(i3_s,i2_s,i3_1,i2_1)
    fault == "c13:0" && (c13_s = c13_1)
    fault == "c13:1" && (c13_s = not(c13_1))
    # -----------------------------------------------------------
    e5_s, e5_1 = op2or(i5_s,i4_s,i5_1,i4_1)
    fault == "e5:0" && (e5_s = e5_1)
    fault == "e5:1" && (e5_s = not(e5_1))
    # -----------------------------------------------------------
    e10_s, e10_1 = op3or(i6_s,i2_s,i0_s,i6_1,i2_1,i0_1)
    fault == "e10:0" && (e10_s = e10_1)
    fault == "e10:1" && (e10_s = not(e10_1))
    # -----------------------------------------------------------
    e12_s, e12_1 = op2or(i6_s,i3_s,i6_1,i3_1)
    fault == "e12:0" && (e12_s = e12_1)
    fault == "e12:1" && (e12_s = not(e12_1))
    # -----------------------------------------------------------
    e18_s, e18_1 = op2or(i6_s,i2_s,i6_1,i2_1)
    fault == "e18:0" && (e18_s = e18_1)
    fault == "e18:1" && (e18_s = not(e18_1))
    # -----------------------------------------------------------
    e24_s, e24_1 = op2or(i7_s,i1_s,i7_1,i1_1)
    fault == "e24:0" && (e24_s = e24_1)
    fault == "e24:1" && (e24_s = not(e24_1))
    # -----------------------------------------------------------
    
    # level 4

    # -----------------------------------------------------------
    a1_s, a1_1 = op2and(s10_s,i0_s,s10_1,i0_1)
    fault == "a1:0" && (a1_s = a1_1)
    fault == "a1:1" && (a1_s = not(a1_1))
    # fault_B == "a1:0" && (a1_1 = not(a1_1))
    # -----------------------------------------------------------
    a2_s, a2_1 = op2and(s15_s,i5_s,s15_1,i5_1)
    fault == "a2:0" && (a2_s = a2_1)
    fault == "a2:1" && (a2_s = not(a2_1))
    # fault_B == "a2:0" && (a2_1 = not(a2_1))
    # -----------------------------------------------------------
    a3_s, a3_1 = op2and(s18_s,ni6_s,s18_1,ni6_1)
    fault == "a3:0" && (a3_s = a3_1)
    fault == "a3:1" && (a3_s = not(a3_1))
    # fault_B == "a3:0" && (a3_1 = not(a3_1))
    # -----------------------------------------------------------
    a4_s, a4_1 = op3or(s20_s,s21_s,s22_s,s20_1,s21_1,s22_1)
    fault == "a4:0" && (a4_s = a4_1)
    fault == "a4:1" && (a4_s = not(a4_1))
    # fault_B == "a4:0" && (a4_1 = not(a4_1))
    # -----------------------------------------------------------
    a5_s, a5_1 = op2and(s24_s,ni7_s,s24_1,ni7_1)
    fault == "a5:0" && (a5_s = a5_1)
    fault == "a5:1" && (a5_s = not(a5_1))
    # fault_B == "a5:0" && (a5_1 = not(a5_1))
    # -----------------------------------------------------------
    a6_s, a6_1 = op2or(s25_s,s26_s,s25_1,s26_1)
    fault == "a6:0" && (a6_s = a6_1)
    fault == "a6:1" && (a6_s = not(a6_1))
    # fault_B == "a6:0" && (a6_1 = not(a6_1))
    # -----------------------------------------------------------
    a7_s, a7_1 = op3or(s27_s,s28_s,s29_s,s27_1,s28_1,s29_1)
    fault == "a7:0" && (a7_s = a7_1)
    fault == "a7:1" && (a7_s = not(a7_1))
    # fault_B == "a7:0" && (a7_1 = not(a7_1))
    # -----------------------------------------------------------
    a9_s, a9_1 = op3and(s31_s,ni5_s,ni2_s,s31_1,ni5_1,ni2_1)
    fault == "a9:0" && (a9_s = a9_1)
    fault == "a9:1" && (a9_s = not(a9_1))
    # fault_B == "a9:0" && (a9_1 = not(a9_1))
    # -----------------------------------------------------------
    b1_s, b1_1 = op2and(s3_s,i2_s,s3_1,i2_1)
    fault == "b1:0" && (b1_s = b1_1)
    fault == "b1:1" && (b1_s = not(b1_1))
    # fault_B == "b1:0" && (b1_1 = not(b1_1))
    # -----------------------------------------------------------
    b3_s, b3_1 = op2and(b2_s,s7_s,b2_1,s7_1)
    fault == "b3:0" && (b3_s = b3_1)
    fault == "b3:1" && (b3_s = not(b3_1))
    # fault_B == "b3:0" && (b3_1 = not(b3_1))
    # -----------------------------------------------------------
    b4_s, b4_1 = op2and(s10_s,ni6_s,s10_1,ni6_1)
    fault == "b4:0" && (b4_s = b4_1)
    fault == "b4:1" && (b4_s = not(b4_1))
    # fault_B == "b4:0" && (b4_1 = not(b4_1))
    # -----------------------------------------------------------
    b5_s, b5_1 = op3or(s12_s,s13_s,s14_s,s12_1,s13_1,s14_1)
    fault == "b5:0" && (b5_s = b5_1)
    fault == "b5:1" && (b5_s = not(b5_1))
    # fault_B == "b5:0" && (b5_1 = not(b5_1))
    # -----------------------------------------------------------
    b6_s, b6_1 = op2and(s15_s,ni5_s,s15_1,ni5_1)
    fault == "b6:0" && (b6_s = b6_1)
    fault == "b6:1" && (b6_s = not(b6_1))
    # fault_B == "b6:0" && (b6_1 = not(b6_1))
    # -----------------------------------------------------------
    b8_s, b8_1 = op2and(s23_s,b7_s,s23_1,b7_1)
    fault == "b8:0" && (b8_s = b8_1)
    fault == "b8:1" && (b8_s = not(b8_1))
    # fault_B == "b8:0" && (b8_1 = not(b8_1))
    # -----------------------------------------------------------
    b9_s, b9_1 = op2or(s25_s,s26_s,s25_1,s26_1)
    fault == "b9:0" && (b9_s = b9_1)
    fault == "b9:1" && (b9_s = not(b9_1))
    # fault_B == "b9:0" && (b9_1 = not(b9_1))
    # -----------------------------------------------------------
    b10_s, b10_1 = op3or(s27_s,s28_s,s29_s,s27_1,s28_1,s29_1)
    fault == "b10:0" && (b10_s = b10_1)
    fault == "b10:1" && (b10_s = not(b10_1))
    # fault_B == "b10:0" && (b10_1 = not(b10_1))
    # -----------------------------------------------------------
    c2_s, c2_1 = op3and(s7_s,ni5_s,ni3_s,s7_1,ni5_1,ni3_1)
    fault == "c2:0" && (c2_s = c2_1)
    fault == "c2:1" && (c2_s = not(c2_1))
    # fault_B == "c2:0" && (c2_1 = not(c2_1))
    # -----------------------------------------------------------
    c3_s, c3_1 = op2and(s11_s,i7_s,s11_1,i7_1)
    fault == "c3:0" && (c3_s = c3_1)
    fault == "c3:1" && (c3_s = not(c3_1))
    # fault_B == "c3:0" && (c3_1 = not(c3_1))
    # -----------------------------------------------------------
    c4_s, c4_1 = op2and(s15_s,ni5_s,s15_1,ni5_1)
    fault == "c4:0" && (c4_s = c4_1)
    fault == "c4:1" && (c4_s = not(c4_1))
    # fault_B == "c4:0" && (c4_1 = not(c4_1))
    # -----------------------------------------------------------
    c6_s, c6_1 = op2and(s23_s,ni5_s,s23_1,ni5_1)
    fault == "c6:0" && (c6_s = c6_1)
    fault == "c6:1" && (c6_s = not(c6_1))
    # fault_B == "c6:0" && (c6_1 = not(c6_1))
    # -----------------------------------------------------------
    c8_s, c8_1 = op2and(s19_s,c5_s,s19_1,c5_1)
    fault == "c8:0" && (c8_s = c8_1)
    fault == "c8:1" && (c8_s = not(c8_1))
    # fault_B == "c8:0" && (c8_1 = not(c8_1))
    # -----------------------------------------------------------
    c14_s, c14_1 = op2and(c13_s,s3_s,c13_1,s3_1)
    fault == "c14:0" && (c14_s = c14_1)
    fault == "c14:1" && (c14_s = not(c14_1))
    # fault_B == "c14:0" && (c14_1 = not(c14_1))
    # -----------------------------------------------------------
    c15_s, c15_1 = op2and(s27_s,i7_s,s27_1,i7_1)
    fault == "c15:0" && (c15_s = c15_1)
    fault == "c15:1" && (c15_s = not(c15_1))
    # fault_B == "c15:0" && (c15_1 = not(c15_1))
    # -----------------------------------------------------------
    d1_s, d1_1 = op2and(s1_s,i2_s,s1_1,i2_1)
    fault == "d1:0" && (d1_s = d1_1)
    fault == "d1:1" && (d1_s = not(d1_1))
    # fault_B == "d1:0" && (d1_1 = not(d1_1))
    # -----------------------------------------------------------
    d2_s, d2_1 = op3and(s3_s,ni3_s,ni2_s,s3_1,ni3_1,ni2_1)
    fault == "d2:0" && (d2_s = d2_1)
    fault == "d2:1" && (d2_s = not(d2_1))
    # fault_B == "d2:0" && (d2_1 = not(d2_1))
    # -----------------------------------------------------------
    d3_s, d3_1 = op2and(s5_s,i0_s,s5_1,i0_1)
    fault == "d3:0" && (d3_s = d3_1)
    fault == "d3:1" && (d3_s = not(d3_1))
    # fault_B == "d3:0" && (d3_1 = not(d3_1))
    # -----------------------------------------------------------
    d5_s, d5_1 = op2and(s9_s,i2_s,s9_1,i2_1)
    fault == "d5:0" && (d5_s = d5_1)
    fault == "d5:1" && (d5_s = not(d5_1))
    # fault_B == "d5:0" && (d5_1 = not(d5_1))
    # -----------------------------------------------------------
    d6_s, d6_1 = op2and(s11_s,ni7_s,s11_1,ni7_1)
    fault == "d6:0" && (d6_s = d6_1)
    fault == "d6:1" && (d6_s = not(d6_1))
    # fault_B == "d6:0" && (d6_1 = not(d6_1))
    # -----------------------------------------------------------
    d7_s, d7_1 = op2and(s13_s,i0_s,s13_1,i0_1)
    fault == "d7:0" && (d7_s = d7_1)
    fault == "d7:1" && (d7_s = not(d7_1))
    # fault_B == "d7:0" && (d7_1 = not(d7_1))
    # -----------------------------------------------------------
    d9_s, d9_1 = op2and(s15_s,ni5_s,s15_1,ni5_1)
    fault == "d9:0" && (d9_s = d9_1)
    fault == "d9:1" && (d9_s = not(d9_1))
    # fault_B == "d9:0" && (d9_1 = not(d9_1))
    # -----------------------------------------------------------
    d10_s, d10_1 = op2and(s17_s,i2_s,s17_1,i2_1)
    fault == "d10:0" && (d10_s = d10_1)
    fault == "d10:1" && (d10_s = not(d10_1))
    # fault_B == "d10:0" && (d10_1 = not(d10_1))
    # -----------------------------------------------------------
    d11_s, d11_1 = op2and(s19_s,ni7_s,s19_1,ni7_1)
    fault == "d11:0" && (d11_s = d11_1)
    fault == "d11:1" && (d11_s = not(d11_1))
    # fault_B == "d11:0" && (d11_1 = not(d11_1))
    # -----------------------------------------------------------
    d12_s, d12_1 = op2and(s21_s,i0_s,s21_1,i0_1)
    fault == "d12:0" && (d12_s = d12_1)
    fault == "d12:1" && (d12_s = not(d12_1))
    # fault_B == "d12:0" && (d12_1 = not(d12_1))
    # -----------------------------------------------------------
    d13_s, d13_1 = op3and(s23_s,ni5_s,ni1_s,s23_1,ni5_1,ni1_1)
    fault == "d13:0" && (d13_s = d13_1)
    fault == "d13:1" && (d13_s = not(d13_1))
    # fault_B == "d13:0" && (d13_1 = not(d13_1))
    # -----------------------------------------------------------
    d14_s, d14_1 = op2and(s25_s,i2_s,s25_1,i2_1)
    fault == "d14:0" && (d14_s = d14_1)
    fault == "d14:1" && (d14_s = not(d14_1))
    # fault_B == "d14:0" && (d14_1 = not(d14_1))
    # -----------------------------------------------------------
    d15_s, d15_1 = op2and(s29_s,i0_s,s29_1,i0_1)
    fault == "d15:0" && (d15_s = d15_1)
    fault == "d15:1" && (d15_s = not(d15_1))
    # fault_B == "d15:0" && (d15_1 = not(d15_1))
    # -----------------------------------------------------------
    d27_s, d27_1 = op2and(s27_s,ni7_s,s27_1,ni7_1)
    fault == "d27:0" && (d27_s = d27_1)
    fault == "d27:1" && (d27_s = not(d27_1))
    # fault_B == "d27:0" && (d27_1 = not(d27_1))
    # -----------------------------------------------------------
    e1_s, e1_1 = op2and(s0_s,i1_s,s0_1,i1_1)
    fault == "e1:0" && (e1_s = e1_1)
    fault == "e1:1" && (e1_s = not(e1_1))
    # fault_B == "e1:0" && (e1_1 = not(e1_1))
    # -----------------------------------------------------------
    e2_s, e2_1 = op2and(s1_s,ni2_s,s1_1,ni2_1)
    fault == "e2:0" && (e2_s = e2_1)
    fault == "e2:1" && (e2_s = not(e2_1))
    # fault_B == "e2:0" && (e2_1 = not(e2_1))
    # -----------------------------------------------------------
    e3_s, e3_1 = op2and(s2_s,i2_s,s2_1,i2_1)
    fault == "e3:0" && (e3_s = e3_1)
    fault == "e3:1" && (e3_s = not(e3_1))
    # fault_B == "e3:0" && (e3_1 = not(e3_1))
    # -----------------------------------------------------------
    e4_s, e4_1 = op3and(s3_s,ni3_s,ni2_s,s3_1,ni3_1,ni2_1)
    fault == "e4:0" && (e4_s = e4_1)
    fault == "e4:1" && (e4_s = not(e4_1))
    # fault_B == "e4:0" && (e4_1 = not(e4_1))
    # -----------------------------------------------------------
    e6_s, e6_1 = op2and(s5_s,ni0_s,s5_1,ni0_1)
    fault == "e6:0" && (e6_s = e6_1)
    fault == "e6:1" && (e6_s = not(e6_1))
    # fault_B == "e6:0" && (e6_1 = not(e6_1))
    # -----------------------------------------------------------
    e7_s, e7_1 = op2and(s6_s,i7_s,s6_1,i7_1)
    fault == "e7:0" && (e7_s = e7_1)
    fault == "e7:1" && (e7_s = not(e7_1))
    # fault_B == "e7:0" && (e7_1 = not(e7_1))
    # -----------------------------------------------------------
    e8_s, e8_1 = op2and(s8_s,i1_s,s8_1,i1_1)
    fault == "e8:0" && (e8_s = e8_1)
    fault == "e8:1" && (e8_s = not(e8_1))
    # fault_B == "e8:0" && (e8_1 = not(e8_1))
    # -----------------------------------------------------------
    e9_s, e9_1 = op2and(s9_s,ni2_s,s9_1,ni2_1)
    fault == "e9:0" && (e9_s = e9_1)
    fault == "e9:1" && (e9_s = not(e9_1))
    # fault_B == "e9:0" && (e9_1 = not(e9_1))
    # -----------------------------------------------------------
    e11_s, e11_1 = op2and(s11_s,ni7_s,s11_1,ni7_1)
    fault == "e11:0" && (e11_s = e11_1)
    fault == "e11:1" && (e11_s = not(e11_1))
    # fault_B == "e11:0" && (e11_1 = not(e11_1))
    # -----------------------------------------------------------
    e13_s, e13_1 = op2and(s13_s,ni0_s,s13_1,ni0_1)
    fault == "e13:0" && (e13_s = e13_1)
    fault == "e13:1" && (e13_s = not(e13_1))
    # fault_B == "e13:0" && (e13_1 = not(e13_1))
    # -----------------------------------------------------------
    e14_s, e14_1 = op2and(s14_s,i7_s,s14_1,i7_1)
    fault == "e14:0" && (e14_s = e14_1)
    fault == "e14:1" && (e14_s = not(e14_1))
    # fault_B == "e14:0" && (e14_1 = not(e14_1))
    # -----------------------------------------------------------
    e15_s, e15_1 = op2and(s15_s,ni5_s,s15_1,ni5_1)
    fault == "e15:0" && (e15_s = e15_1)
    fault == "e15:1" && (e15_s = not(e15_1))
    # fault_B == "e15:0" && (e15_1 = not(e15_1))
    # -----------------------------------------------------------
    e16_s, e16_1 = op2and(s16_s,i1_s,s16_1,i1_1)
    fault == "e16:0" && (e16_s = e16_1)
    fault == "e16:1" && (e16_s = not(e16_1))
    # fault_B == "e16:0" && (e16_1 = not(e16_1))
    # -----------------------------------------------------------
    e17_s, e17_1 = op2and(s17_s,ni2_s,s17_1,ni2_1)
    fault == "e17:0" && (e17_s = e17_1)
    fault == "e17:1" && (e17_s = not(e17_1))
    # fault_B == "e17:0" && (e17_1 = not(e17_1))
    # -----------------------------------------------------------
    e19_s, e19_1 = op2and(s19_s,ni7_s,s19_1,ni7_1)
    fault == "e19:0" && (e19_s = e19_1)
    fault == "e19:1" && (e19_s = not(e19_1))
    # fault_B == "e19:0" && (e19_1 = not(e19_1))
    # -----------------------------------------------------------
    e20_s, e20_1 = op2and(s20_s,e12_s,s20_1,e12_1)
    fault == "e20:0" && (e20_s = e20_1)
    fault == "e20:1" && (e20_s = not(e20_1))
    # fault_B == "e20:0" && (e20_1 = not(e20_1))
    # -----------------------------------------------------------
    e21_s, e21_1 = op2and(s21_s,ni0_s,s21_1,ni0_1)
    fault == "e21:0" && (e21_s = e21_1)
    fault == "e21:1" && (e21_s = not(e21_1))
    # fault_B == "e21:0" && (e21_1 = not(e21_1))
    # -----------------------------------------------------------
    e22_s, e22_1 = op2and(s22_s,i7_s,s22_1,i7_1)
    fault == "e22:0" && (e22_s = e22_1)
    fault == "e22:1" && (e22_s = not(e22_1))
    # fault_B == "e22:0" && (e22_1 = not(e22_1))
    # -----------------------------------------------------------
    e23_s, e23_1 = op2and(s23_s,ni5_s,s23_1,ni5_1)
    fault == "e23:0" && (e23_s = e23_1)
    fault == "e23:1" && (e23_s = not(e23_1))
    # fault_B == "e23:0" && (e23_1 = not(e23_1))
    # -----------------------------------------------------------
    e25_s, e25_1 = op2and(s25_s,ni2_s,s25_1,ni2_1)
    fault == "e25:0" && (e25_s = e25_1)
    fault == "e25:1" && (e25_s = not(e25_1))
    # fault_B == "e25:0" && (e25_1 = not(e25_1))
    # -----------------------------------------------------------
    e26_s, e26_1 = op2and(s26_s,i2_s,s26_1,i2_1)
    fault == "e26:0" && (e26_s = e26_1)
    fault == "e26:1" && (e26_s = not(e26_1))
    # fault_B == "e26:0" && (e26_1 = not(e26_1))
    # -----------------------------------------------------------
    e27_s, e27_1 = op2and(s27_s,ni7_s,s27_1,ni7_1)
    fault == "e27:0" && (e27_s = e27_1)
    fault == "e27:1" && (e27_s = not(e27_1))
    # fault_B == "e27:0" && (e27_1 = not(e27_1))
    # -----------------------------------------------------------
    e28_s, e28_1 = op2and(s28_s,e12_s,s28_1,e12_1)
    fault == "e28:0" && (e28_s = e28_1)
    fault == "e28:1" && (e28_s = not(e28_1))
    # fault_B == "e28:0" && (e28_1 = not(e28_1))
    # -----------------------------------------------------------
    e29_s, e29_1 = op2and(s29_s,ni0_s,s29_1,ni0_1)
    fault == "e29:0" && (e29_s = e29_1)
    fault == "e29:1" && (e29_s = not(e29_1))
    # fault_B == "e29:0" && (e29_1 = not(e29_1))
    # -----------------------------------------------------------
    e30_s, e30_1 = op2and(s30_s,i7_s,s30_1,i7_1)
    fault == "e30:0" && (e30_s = e30_1)
    fault == "e30:1" && (e30_s = not(e30_1))
    # fault_B == "e30:0" && (e30_1 = not(e30_1))
    # -----------------------------------------------------------
    e31_s, e31_1 = op2and(s4_s,e5_s,s4_1,e5_1)
    fault == "e31:0" && (e31_s = e31_1)
    fault == "e31:1" && (e31_s = not(e31_1))
    # fault_B == "e31:0" && (e31_1 = not(e31_1))
    # -----------------------------------------------------------
    e32_s, e32_1 = op2and(s10_s,e10_s,s10_1,e10_1)
    fault == "e32:0" && (e32_s = e32_1)
    fault == "e32:1" && (e32_s = not(e32_1))
    # fault_B == "e32:0" && (e32_1 = not(e32_1))
    # -----------------------------------------------------------
    e33_s, e33_1 = op2and(s12_s,e12_s,s12_1,e12_1)
    fault == "e33:0" && (e33_s = e33_1)
    fault == "e33:1" && (e33_s = not(e33_1))
    # fault_B == "e33:0" && (e33_1 = not(e33_1))
    # -----------------------------------------------------------
    e34_s, e34_1 = op2and(s18_s,e18_s,s18_1,e18_1)
    fault == "e34:0" && (e34_s = e34_1)
    fault == "e34:1" && (e34_s = not(e34_1))
    # fault_B == "e34:0" && (e34_1 = not(e34_1))
    # -----------------------------------------------------------
    e35_s, e35_1 = op2and(s24_s,e24_s,s24_1,e24_1)
    fault == "e35:0" && (e35_s = e35_1)
    fault == "e35:1" && (e35_s = not(e35_1))
    # fault_B == "e35:0" && (e35_1 = not(e35_1))
    # -----------------------------------------------------------
    f1_s, f1_1 = op2and(s12_s,i5_s,s12_1,i5_1)
    fault == "f1:0" && (f1_s = f1_1)
    fault == "f1:1" && (f1_s = not(f1_1))
    # fault_B == "f1:0" && (f1_1 = not(f1_1))
    # -----------------------------------------------------------
    f2_s, f2_1 = op2and(s27_s,i4_s,s27_1,i4_1)
    fault == "f2:0" && (f2_s = f2_1)
    fault == "f2:1" && (f2_s = not(f2_1))
    # fault_B == "f2:0" && (f2_1 = not(f2_1))
    # -----------------------------------------------------------
    f3_s, f3_1 = op2and(s15_s,i0_s,s15_1,i0_1)
    fault == "f3:0" && (f3_s = f3_1)
    fault == "f3:1" && (f3_s = not(f3_1))
    # fault_B == "f3:0" && (f3_1 = not(f3_1))
    # -----------------------------------------------------------
    f4_s, f4_1 = op2and(s27_s,i2_s,s27_1,i2_1)
    fault == "f4:0" && (f4_s = f4_1)
    fault == "f4:1" && (f4_s = not(f4_1))
    # fault_B == "f4:0" && (f4_1 = not(f4_1))
    # -----------------------------------------------------------
    f5_s, f5_1 = op2and(s0_s,i7_s,s0_1,i7_1)
    fault == "f5:0" && (f5_s = f5_1)
    fault == "f5:1" && (f5_s = not(f5_1))
    # fault_B == "f5:0" && (f5_1 = not(f5_1))
    # -----------------------------------------------------------
    f6_s, f6_1 = op2and(s27_s,i1_s,s27_1,i1_1)
    fault == "f6:0" && (f6_s = f6_1)
    fault == "f6:1" && (f6_s = not(f6_1))
    # fault_B == "f6:0" && (f6_1 = not(f6_1))
    # -----------------------------------------------------------

    # level 5

    # -----------------------------------------------------------
    a8_s,a8_1 = op2or(a7_s,s30_s,a7_1,s30_1)
    fault == "a8:0" && (a8_s = a8_1)
    fault == "a8:1" && (a8_s = not(a8_1))
    # fault_B == "a8:0" && (a8_1 = not(a8_1))
    # -----------------------------------------------------------
    a10_s,a10_1 = op3or(a1_s,a2_s,s16_s,a1_1,a2_1,s16_1)
    fault == "a10:0" && (a10_s = a10_1)
    fault == "a10:1" && (a10_s = not(a10_1))
    # fault_B == "a10:0" && (a10_1 = not(a10_1))
    # -----------------------------------------------------------
    b11_s,b11_1 = op2or(b10_s,s30_s,b10_1,s30_1)
    fault == "b11:0" && (b11_s = b11_1)
    fault == "b11:1" && (b11_s = not(b11_1))
    # fault_B == "b11:0" && (b11_1 = not(b11_1))
    # -----------------------------------------------------------
    b12_s,b12_1 = op3or(b1_s,b3_s,s8_s,b1_1,b3_1,s8_1)
    fault == "b12:0" && (b12_s = b12_1)
    fault == "b12:1" && (b12_s = not(b12_1))
    # fault_B == "b12:0" && (b12_1 = not(b12_1))
    # -----------------------------------------------------------
    b13_s,b13_1 = op3or(s9_s,b4_s,s11_s,s9_1,b4_1,s11_1)
    fault == "b13:0" && (b13_s = b13_1)
    fault == "b13:1" && (b13_s = not(b13_1))
    # fault_B == "b13:0" && (b13_1 = not(b13_1))
    # -----------------------------------------------------------
    c1_s,c1_1 = op3or(c14_s,s4_s,s5_s,c14_1,s4_1,s5_1)
    fault == "c1:0" && (c1_s = c1_1)
    fault == "c1:1" && (c1_s = not(c1_1))
    # fault_B == "c1:0" && (c1_1 = not(c1_1))
    # -----------------------------------------------------------
    c7_s,c7_1 = op2and(c2_s,ni2_s,c2_1,ni2_1)
    fault == "c7:0" && (c7_s = c7_1)
    fault == "c7:1" && (c7_s = not(c7_1))
    # fault_B == "c7:0" && (c7_1 = not(c7_1))
    # -----------------------------------------------------------
    c16_s,c16_1 = op3or(c15_s,s28_s,s29_s,c15_1,s28_1,s29_1)
    fault == "c16:0" && (c16_s = c16_1)
    fault == "c16:1" && (c16_s = not(c16_1))
    # fault_B == "c16:0" && (c16_1 = not(c16_1))
    # -----------------------------------------------------------
    d17_s,d17_1 = op3or(d1_s,d2_s,d3_s,d1_1,d2_1,d3_1)
    fault == "d17:0" && (d17_s = d17_1)
    fault == "d17:1" && (d17_s = not(d17_1))
    # fault_B == "d17:0" && (d17_1 = not(d17_1))
    # -----------------------------------------------------------
    d19_s,d19_1 = op3or(s10_s,d6_s,d7_s,s10_1,d6_1,d7_1)
    fault == "d19:0" && (d19_s = d19_1)
    fault == "d19:1" && (d19_s = not(d19_1))
    # fault_B == "d19:0" && (d19_1 = not(d19_1))
    # -----------------------------------------------------------
    d20_s,d20_1 = op3or(s14_s,d9_s,d10_s,s14_1,d9_1,d10_1)
    fault == "d20:0" && (d20_s = d20_1)
    fault == "d20:1" && (d20_s = not(d20_1))
    # fault_B == "d20:0" && (d20_1 = not(d20_1))
    # -----------------------------------------------------------
    d21_s,d21_1 = op3or(s18_s,d11_s,d12_s,s18_1,d11_1,d12_1)
    fault == "d21:0" && (d21_s = d21_1)
    fault == "d21:1" && (d21_s = not(d21_1))
    # fault_B == "d21:0" && (d21_1 = not(d21_1))
    # -----------------------------------------------------------
    d22_s,d22_1 = op3or(s22_s,d13_s,d14_s,s22_1,d13_1,d14_1)
    fault == "d22:0" && (d22_s = d22_1)
    fault == "d22:1" && (d22_s = not(d22_1))
    # fault_B == "d22:0" && (d22_1 = not(d22_1))
    # -----------------------------------------------------------
    d23_s,d23_1 = op3or(s26_s,d15_s,s30_s,s26_1,d15_1,s30_1)
    fault == "d23:0" && (d23_s = d23_1)
    fault == "d23:1" && (d23_s = not(d23_1))
    # fault_B == "d23:0" && (d23_1 = not(d23_1))
    # -----------------------------------------------------------
    e36_s,e36_1 = op3or(e1_s,e2_s,e3_s,e1_1,e2_1,e3_1)
    fault == "e36:0" && (e36_s = e36_1)
    fault == "e36:1" && (e36_s = not(e36_1))
    # fault_B == "e36:0" && (e36_1 = not(e36_1))
    # -----------------------------------------------------------
    e37_s,e37_1 = op3or(e4_s,e31_s,e6_s,e4_1,e31_1,e6_1)
    fault == "e37:0" && (e37_s = e37_1)
    fault == "e37:1" && (e37_s = not(e37_1))
    # fault_B == "e37:0" && (e37_1 = not(e37_1))
    # -----------------------------------------------------------
    e39_s,e39_1 = op2or(e9_s,e32_s,e9_1,e32_1)
    fault == "e39:0" && (e39_s = e39_1)
    fault == "e39:1" && (e39_s = not(e39_1))
    # fault_B == "e39:0" && (e39_1 = not(e39_1))
    # -----------------------------------------------------------
    e40_s,e40_1 = op3or(e11_s,e33_s,e13_s,e11_1,e33_1,e13_1)
    fault == "e40:0" && (e40_s = e40_1)
    fault == "e40:1" && (e40_s = not(e40_1))
    # fault_B == "e40:0" && (e40_1 = not(e40_1))
    # -----------------------------------------------------------
    e41_s,e41_1 = op3or(e14_s,e15_s,e16_s,e14_1,e15_1,e16_1)
    fault == "e41:0" && (e41_s = e41_1)
    fault == "e41:1" && (e41_s = not(e41_1))
    # fault_B == "e41:0" && (e41_1 = not(e41_1))
    # -----------------------------------------------------------
    e42_s, e42_1 = op3or(e17_s,e34_s,e19_s,e17_1,e34_1,e19_1)
    fault == "e42:0" && (e42_s = e42_1)
    fault == "e42:1" && (e42_s = not(e42_1))
    # fault_B == "e42:0" && (e42_1 = not(e42_1))
    # -----------------------------------------------------------    
    e43_s,e43_1 = op3or(e20_s,e30_s,a9_s,e20_1,e30_1,a9_1)
    fault == "e43:0" && (e43_s = e43_1)
    fault == "e43:1" && (e43_s = not(e43_1))
    # fault_B == "e43:0" && (e43_1 = not(e43_1))
    # -----------------------------------------------------------
    e44_s,e44_1 = op3or(e21_s,e22_s,e23_s,e21_1,e22_1,e23_1)
    fault == "e44:0" && (e44_s = e44_1)
    fault == "e44:1" && (e44_s = not(e44_1))
    # fault_B == "e44:0" && (e44_1 = not(e44_1))
    # -----------------------------------------------------------
    e45_s,e45_1 = op3or(e35_s,e25_s,e26_s,e35_1,e25_1,e26_1)
    fault == "e45:0" && (e45_s = e45_1)
    fault == "e45:1" && (e45_s = not(e45_1))
    # fault_B == "e45:0" && (e45_1 = not(e45_1))
    # -----------------------------------------------------------
    e46_s,e46_1 = op3or(e27_s,e28_s,e29_s,e27_1,e28_1,e29_1)
    fault == "e46:0" && (e46_s = e46_1)
    fault == "e46:1" && (e46_s = not(e46_1))
    # fault_B == "e46:0" && (e46_1 = not(e46_1))
    # -----------------------------------------------------------
    out4_s,out4_1 = op2or(f1_s,f2_s,f1_1,f2_1)
    fault == "out4:0" && (out4_s = out4_1)
    fault == "out4:1" && (out4_s = not(out4_1))
    # fault_B == "out4:0" && (out4_1 = not(out4_1))
    # -----------------------------------------------------------
    out2_s,out2_1 = op2or(f3_s,f4_s,f3_1,f4_1)
    fault == "out2:0" && (out2_s = out2_1)
    fault == "out2:1" && (out2_s = not(out2_1))
    # fault_B == "out2:0" && (out2_1 = not(out2_1))
    # -----------------------------------------------------------
    out1_s,out1_1 = op2or(f5_s,f6_s,f5_1,f6_1)
    fault == "out1:0" && (out1_s = out1_1)
    fault == "out1:1" && (out1_s = not(out1_1))
    # fault_B == "out1:0" && (out1_1 = not(out1_1))
    # -----------------------------------------------------------

    # level 6

    # -----------------------------------------------------------
    a11_s,a11_1 = op3or(a10_s,s17_s,a3_s,a10_1,s17_1,a3_1)
    fault == "a11:0" && (a11_s = a11_1)
    fault == "a11:1" && (a11_s = not(a11_1))
    # fault_B == "a11:0" && (a11_1 = not(a11_1))
    # -----------------------------------------------------------
    b14_s,b14_1 = op3or(b12_s,b13_s,b5_s,b12_1,b13_1,b5_1)
    fault == "b14:0" && (b14_s = b14_1)
    fault == "b14:1" && (b14_s = not(b14_1))
    # fault_B == "b14:0" && (b14_1 = not(b14_1))
    # -----------------------------------------------------------
    c9_s,c9_1 = op3or(c1_s,s6_s,c7_s,c1_1,s6_1,c7_1)
    fault == "c9:0" && (c9_s = c9_1)
    fault == "c9:1" && (c9_s = not(c9_1))
    # fault_B == "c9:0" && (c9_1 = not(c9_1))
    # -----------------------------------------------------------
    c17_s,c17_1 = op2or(c16_s,s30_s,c16_1,s30_1)
    fault == "c17:0" && (c17_s = c17_1)
    fault == "c17:1" && (c17_s = not(c17_1))
    # fault_B == "c17:0" && (c17_1 = not(c17_1))
    # -----------------------------------------------------------
    d18_s,d18_1 = op3or(s6_s,c7_s,d5_s,s6_1,c7_1,d5_1)
    fault == "d18:0" && (d18_s = d18_1)
    fault == "d18:1" && (d18_s = not(d18_1))
    # fault_B == "d18:0" && (d18_1 = not(d18_1))
    # -----------------------------------------------------------
    d28_s,d28_1 = op2or(d23_s,d27_s,d23_1,d27_1)
    fault == "d28:0" && (d28_s = d28_1)
    fault == "d28:1" && (d28_s = not(d28_1))
    # fault_B == "d28:0" && (d28_1 = not(d28_1))
    # -----------------------------------------------------------
    e38_s,e38_1 = op3or(e7_s,c7_s,e8_s,e7_1,c7_1,e8_1)
    fault == "e38:0" && (e38_s = e38_1)
    fault == "e38:1" && (e38_s = not(e38_1))
    # fault_B == "e38:0" && (e38_1 = not(e38_1))
    # -----------------------------------------------------------
    e47_s,e47_1 = op3or(e36_s,e37_s,e44_s,e36_1,e37_1,e44_1)
    fault == "e47:0" && (e47_s = e47_1)
    fault == "e47:1" && (e47_s = not(e47_1))
    # fault_B == "e47:0" && (e47_1 = not(e47_1))
    # -----------------------------------------------------------
    e49_s,e49_1 = op3or(e39_s,e40_s,e41_s,e39_1,e40_1,e41_1)
    fault == "e49:0" && (e49_s = e49_1)
    fault == "e49:1" && (e49_s = not(e49_1))
    # fault_B == "e49:0" && (e49_1 = not(e49_1))
    # -----------------------------------------------------------

    # level 7

    # -----------------------------------------------------------
    a12_s,a12_1 = op3or(a11_s,s19_s,a4_s,a11_1,s19_1,a4_1)
    fault == "a12:0" && (a12_s = a12_1)
    fault == "a12:1" && (a12_s = not(a12_1))
    # fault_B == "a12:0" && (a12_1 = not(a12_1))
    # -----------------------------------------------------------
    b15_s,b15_1 = op3or(b14_s,b6_s,b8_s,b14_1,b6_1,b8_1)
    fault == "b15:0" && (b15_s = b15_1)
    fault == "b15:1" && (b15_s = not(b15_1))
    # fault_B == "b15:0" && (b15_1 = not(b15_1))
    # -----------------------------------------------------------
    c10_s,c10_1 = op3or(c9_s,c3_s,b5_s,c9_1,c3_1,b5_1)
    fault == "c10:0" && (c10_s = c10_1)
    fault == "c10:1" && (c10_s = not(c10_1))
    # fault_B == "c10:0" && (c10_1 = not(c10_1))
    # -----------------------------------------------------------
    d24_s,d24_1 = op3or(d17_s,d18_s,d19_s,d17_1,d18_1,d19_1)
    fault == "d24:0" && (d24_s = d24_1)
    fault == "d24:1" && (d24_s = not(d24_1))
    # fault_B == "d24:0" && (d24_1 = not(d24_1))
    # -----------------------------------------------------------
    e48_s,e48_1 = op3or(e45_s,e46_s,e38_s,e45_1,e46_1,e38_1)
    fault == "e48:0" && (e48_s = e48_1)
    fault == "e48:1" && (e48_s = not(e48_1))
    # fault_B == "e48:0" && (e48_1 = not(e48_1))
    # -----------------------------------------------------------

    # level 8

    # -----------------------------------------------------------
    a13_s,a13_1 = op3or(a12_s,s23_s,a5_s,a12_1,s23_1,a5_1)
    fault == "a13:0" && (a13_s = a13_1)
    fault == "a13:1" && (a13_s = not(a13_1))
    # fault_B == "a13:0" && (a13_1 = not(a13_1))
    # -----------------------------------------------------------
    b16_s,b16_1 = op3or(b15_s,s24_s,b9_s,b15_1,s24_1,b9_1)
    fault == "b16:0" && (b16_s = b16_1)
    fault == "b16:1" && (b16_s = not(b16_1))
    # fault_B == "b16:0" && (b16_1 = not(b16_1))
    # -----------------------------------------------------------
    c11_s,c11_1 = op3or(c10_s,c4_s,c8_s,c10_1,c4_1,c8_1)
    fault == "c11:0" && (c11_s = c11_1)
    fault == "c11:1" && (c11_s = not(c11_1))
    # fault_B == "c11:0" && (c11_1 = not(c11_1))
    # -----------------------------------------------------------
    d25_s,d25_1 = op3or(d24_s,d20_s,d21_s,d24_1,d20_1,d21_1)
    fault == "d25:0" && (d25_s = d25_1)
    fault == "d25:1" && (d25_s = not(d25_1))
    # fault_B == "d25:0" && (d25_1 = not(d25_1))
    # -----------------------------------------------------------
    e50_s,e50_1 = op3or(e47_s,e48_s,e49_s,e47_1,e48_1,e49_1)
    fault == "e50:0" && (e50_s = e50_1)
    fault == "e50:1" && (e50_s = not(e50_1))
    # fault_B == "e50:0" && (e50_1 = not(e50_1))
    # -----------------------------------------------------------

    # level 9

    # -----------------------------------------------------------
    ns1_s,ns1_1 = op3or(e50_s,e42_s,e43_s,e50_1,e42_1,e43_1)
    fault == "ns1:0" && (ns1_s = ns1_1)
    fault == "ns1:1" && (ns1_s = not(ns1_1))
    # fault_B == "ns1:0" && (ns1_1 = not(ns1_1))
    # -----------------------------------------------------------
    ns8_s,ns8_1 = op3or(b16_s,b11_s,a9_s,b16_1,b11_1,a9_1)
    fault == "ns8:0" && (ns8_s = ns8_1)
    fault == "ns8:1" && (ns8_s = not(ns8_1))
    # fault_B == "ns8:0" && (ns8_1 = not(ns8_1))
    # -----------------------------------------------------------
    a14_s,a14_1 = op3or(a13_s,a6_s,a8_s,a13_1,a6_1,a8_1)
    fault == "a14:0" && (a14_s = a14_1)
    fault == "a14:1" && (a14_s = not(a14_1))
    # fault_B == "a14:0" && (a14_1 = not(a14_1))
    # -----------------------------------------------------------
    c12_s,c12_1 = op3or(c11_s,a4_s,c6_s,c11_1,a4_1,c6_1)
    fault == "c12:0" && (c12_s = c12_1)
    fault == "c12:1" && (c12_s = not(c12_1))
    # fault_B == "c12:0" && (c12_1 = not(c12_1))
    # -----------------------------------------------------------
    d26_s,d26_1 = op3or(d25_s,d22_s,d28_s,d25_1,d22_1,d28_1)
    fault == "d26:0" && (d26_s = d26_1)
    fault == "d26:1" && (d26_s = not(d26_1))
    # fault_B == "d26:0" && (d26_1 = not(d26_1))
    # -----------------------------------------------------------

    # level 10

    # -----------------------------------------------------------
    ns2_s,ns2_1 = op3or(d26_s,s2_s,a9_s,d26_1,s2_1,a9_1)
    fault == "ns2:0" && (ns2_s = ns2_1)
    fault == "ns2:1" && (ns2_s = not(ns2_1))
    # fault_B == "ns2:0" && (ns2_1 = not(ns2_1))
    # -----------------------------------------------------------
    ns4_s,ns4_1 = op3or(c12_s,c17_s,a9_s,c12_1,c17_1,a9_1)
    fault == "ns4:0" && (ns4_s = ns4_1)
    fault == "ns4:1" && (ns4_s = not(ns4_1))
    # fault_B == "ns4:0" && (ns4_1 = not(ns4_1))
    # -----------------------------------------------------------
    ns16_s,ns16_1 = op2or(a14_s,a9_s,a14_1,a9_1)
    fault == "ns16:0" && (ns16_s = ns16_1)
    fault == "ns16:1" && (ns16_s = not(ns16_1))
    # fault_B == "ns16:0" && (ns16_1 = not(ns16_1))
    # -----------------------------------------------------------

    return (out4_s, out2_s, out1_s, ns16_s, ns8_s, ns4_s, ns2_s, ns1_s)

end # of ps2ns()


# ns2dict ==============================================================
# ======================================================================

function ns2dict(ns16_s, ns8_s, ns4_s, ns2_s, ns1_s)

    function decomp(line)
        
        # receives an ns line, decomposes it into constituent S/Is,
        # maps the S/Is into dictionary form (s, i) => ns, where ns
        # is a next state that follows (s, i), gathers the objects
        # into one dictionary for each ns_line, which may be empty.
        # The right-most column is, in effect, a state-transition
        # table for the circuit Large.
        
        m = Dict()
        
        and(line, and(s0, i0)) != null && (m[(0, 0)] = 0)
        and(line, and(s0, i1)) != null && (m[(0, 1)] = 1)
        and(line, and(s0, i2)) != null && (m[(0, 2)] = 0)
        and(line, and(s0, i3)) != null && (m[(0, 3)] = 0)
        and(line, and(s0, i4)) != null && (m[(0, 4)] = 0)
        and(line, and(s0, i5)) != null && (m[(0, 5)] = 0)
        and(line, and(s0, i6)) != null && (m[(0, 6)] = 0)
        and(line, and(s0, i7)) != null && (m[(0, 7)] = 0)
        and(line, and(s1, i0)) != null && (m[(1, 0)] = 1)
        and(line, and(s1, i1)) != null && (m[(1, 1)] = 1)
        and(line, and(s1, i2)) != null && (m[(1, 2)] = 2)
        and(line, and(s1, i3)) != null && (m[(1, 3)] = 1)
        and(line, and(s1, i4)) != null && (m[(1, 4)] = 1)
        and(line, and(s1, i5)) != null && (m[(1, 5)] = 1)
        and(line, and(s1, i6)) != null && (m[(1, 6)] = 1)
        and(line, and(s1, i7)) != null && (m[(1, 7)] = 1)
        and(line, and(s2, i0)) != null && (m[(2, 0)] = 2)
        and(line, and(s2, i1)) != null && (m[(2, 1)] = 2)
        and(line, and(s2, i2)) != null && (m[(2, 2)] = 3)
        and(line, and(s2, i3)) != null && (m[(2, 3)] = 2)
        and(line, and(s2, i4)) != null && (m[(2, 4)] = 2)
        and(line, and(s2, i5)) != null && (m[(2, 5)] = 2)
        and(line, and(s2, i6)) != null && (m[(2, 6)] = 2)
        and(line, and(s2, i7)) != null && (m[(2, 7)] = 2)
        and(line, and(s3, i0)) != null && (m[(3, 0)] = 3)
        and(line, and(s3, i1)) != null && (m[(3, 1)] = 3)
        and(line, and(s3, i2)) != null && (m[(3, 2)] = 12)
        and(line, and(s3, i3)) != null && (m[(3, 3)] = 4)
        and(line, and(s3, i4)) != null && (m[(3, 4)] = 3)
        and(line, and(s3, i5)) != null && (m[(3, 5)] = 3)
        and(line, and(s3, i6)) != null && (m[(3, 6)] = 3)
        and(line, and(s3, i7)) != null && (m[(3, 7)] = 3)
        and(line, and(s4, i0)) != null && (m[(4, 0)] = 4)
        and(line, and(s4, i1)) != null && (m[(4, 1)] = 4)
        and(line, and(s4, i2)) != null && (m[(4, 2)] = 4)
        and(line, and(s4, i3)) != null && (m[(4, 3)] = 4)
        and(line, and(s4, i4)) != null && (m[(4, 4)] = 5)
        and(line, and(s4, i5)) != null && (m[(4, 5)] = 5)
        and(line, and(s4, i6)) != null && (m[(4, 6)] = 4)
        and(line, and(s4, i7)) != null && (m[(4, 7)] = 4)
        and(line, and(s5, i0)) != null && (m[(5, 0)] = 6)
        and(line, and(s5, i1)) != null && (m[(5, 1)] = 5)
        and(line, and(s5, i2)) != null && (m[(5, 2)] = 5)
        and(line, and(s5, i3)) != null && (m[(5, 3)] = 5)
        and(line, and(s5, i4)) != null && (m[(5, 4)] = 5)
        and(line, and(s5, i5)) != null && (m[(5, 5)] = 5)
        and(line, and(s5, i6)) != null && (m[(5, 6)] = 5)
        and(line, and(s5, i7)) != null && (m[(5, 7)] = 5)
        and(line, and(s6, i0)) != null && (m[(6, 0)] = 6)
        and(line, and(s6, i1)) != null && (m[(6, 1)] = 6)
        and(line, and(s6, i2)) != null && (m[(6, 2)] = 6)
        and(line, and(s6, i3)) != null && (m[(6, 3)] = 6)
        and(line, and(s6, i4)) != null && (m[(6, 4)] = 6)
        and(line, and(s6, i5)) != null && (m[(6, 5)] = 6)
        and(line, and(s6, i6)) != null && (m[(6, 6)] = 6)
        and(line, and(s6, i7)) != null && (m[(6, 7)] = 7)
        and(line, and(s7, i0)) != null && (m[(7, 0)] = 7)
        and(line, and(s7, i1)) != null && (m[(7, 1)] = 7)
        and(line, and(s7, i2)) != null && (m[(7, 2)] = 8)
        and(line, and(s7, i3)) != null && (m[(7, 3)] = 8)
        and(line, and(s7, i4)) != null && (m[(7, 4)] = 7)
        and(line, and(s7, i5)) != null && (m[(7, 5)] = 8)
        and(line, and(s7, i6)) != null && (m[(7, 6)] = 7)
        and(line, and(s7, i7)) != null && (m[(7, 7)] = 7)
        and(line, and(s8, i0)) != null && (m[(8, 0)] = 8)
        and(line, and(s8, i1)) != null && (m[(8, 1)] = 9)
        and(line, and(s8, i2)) != null && (m[(8, 2)] = 8)
        and(line, and(s8, i3)) != null && (m[(8, 3)] = 8)
        and(line, and(s8, i4)) != null && (m[(8, 4)] = 8)
        and(line, and(s8, i5)) != null && (m[(8, 5)] = 8)
        and(line, and(s8, i6)) != null && (m[(8, 6)] = 8)
        and(line, and(s8, i7)) != null && (m[(8, 7)] = 8)
        and(line, and(s9, i0)) != null && (m[(9, 0)] = 9)
        and(line, and(s9, i1)) != null && (m[(9, 1)] = 9)
        and(line, and(s9, i2)) != null && (m[(9, 2)] = 10)
        and(line, and(s9, i3)) != null && (m[(9, 3)] = 9)
        and(line, and(s9, i4)) != null && (m[(9, 4)] = 9)
        and(line, and(s9, i5)) != null && (m[(9, 5)] = 9)
        and(line, and(s9, i6)) != null && (m[(9, 6)] = 9)
        and(line, and(s9, i7)) != null && (m[(9, 7)] = 9)
        and(line, and(s10, i0)) != null && (m[(10, 0)] = 27)
        and(line, and(s10, i1)) != null && (m[(10, 1)] = 10)
        and(line, and(s10, i2)) != null && (m[(10, 2)] = 11)
        and(line, and(s10, i3)) != null && (m[(10, 3)] = 10)
        and(line, and(s10, i4)) != null && (m[(10, 4)] = 10)
        and(line, and(s10, i5)) != null && (m[(10, 5)] = 10)
        and(line, and(s10, i6)) != null && (m[(10, 6)] = 3)
        and(line, and(s10, i7)) != null && (m[(10, 7)] = 10)
        and(line, and(s11, i0)) != null && (m[(11, 0)] = 11)
        and(line, and(s11, i1)) != null && (m[(11, 1)] = 11)
        and(line, and(s11, i2)) != null && (m[(11, 2)] = 11)
        and(line, and(s11, i3)) != null && (m[(11, 3)] = 11)
        and(line, and(s11, i4)) != null && (m[(11, 4)] = 11)
        and(line, and(s11, i5)) != null && (m[(11, 5)] = 11)
        and(line, and(s11, i6)) != null && (m[(11, 6)] = 11)
        and(line, and(s11, i7)) != null && (m[(11, 7)] = 12)
        and(line, and(s12, i0)) != null && (m[(12, 0)] = 12)
        and(line, and(s12, i1)) != null && (m[(12, 1)] = 12)
        and(line, and(s12, i2)) != null && (m[(12, 2)] = 12)
        and(line, and(s12, i3)) != null && (m[(12, 3)] = 13)
        and(line, and(s12, i4)) != null && (m[(12, 4)] = 12)
        and(line, and(s12, i5)) != null && (m[(12, 5)] = 12)
        and(line, and(s12, i6)) != null && (m[(12, 6)] = 13)
        and(line, and(s12, i7)) != null && (m[(12, 7)] = 12)
        and(line, and(s13, i0)) != null && (m[(13, 0)] = 14)
        and(line, and(s13, i1)) != null && (m[(13, 1)] = 13)
        and(line, and(s13, i2)) != null && (m[(13, 2)] = 13)
        and(line, and(s13, i3)) != null && (m[(13, 3)] = 13)
        and(line, and(s13, i4)) != null && (m[(13, 4)] = 13)
        and(line, and(s13, i5)) != null && (m[(13, 5)] = 13)
        and(line, and(s13, i6)) != null && (m[(13, 6)] = 13)
        and(line, and(s13, i7)) != null && (m[(13, 7)] = 13)
        and(line, and(s14, i0)) != null && (m[(14, 0)] = 14)
        and(line, and(s14, i1)) != null && (m[(14, 1)] = 14)
        and(line, and(s14, i2)) != null && (m[(14, 2)] = 14)
        and(line, and(s14, i3)) != null && (m[(14, 3)] = 14)
        and(line, and(s14, i4)) != null && (m[(14, 4)] = 14)
        and(line, and(s14, i5)) != null && (m[(14, 5)] = 14)
        and(line, and(s14, i6)) != null && (m[(14, 6)] = 14)
        and(line, and(s14, i7)) != null && (m[(14, 7)] = 15)
        and(line, and(s15, i0)) != null && (m[(15, 0)] = 15)
        and(line, and(s15, i1)) != null && (m[(15, 1)] = 15)
        and(line, and(s15, i2)) != null && (m[(15, 2)] = 15)
        and(line, and(s15, i3)) != null && (m[(15, 3)] = 15)
        and(line, and(s15, i4)) != null && (m[(15, 4)] = 15)
        and(line, and(s15, i5)) != null && (m[(15, 5)] = 16)
        and(line, and(s15, i6)) != null && (m[(15, 6)] = 15)
        and(line, and(s15, i7)) != null && (m[(15, 7)] = 15)
        and(line, and(s16, i0)) != null && (m[(16, 0)] = 16)
        and(line, and(s16, i1)) != null && (m[(16, 1)] = 17)
        and(line, and(s16, i2)) != null && (m[(16, 2)] = 16)
        and(line, and(s16, i3)) != null && (m[(16, 3)] = 16)
        and(line, and(s16, i4)) != null && (m[(16, 4)] = 16)
        and(line, and(s16, i5)) != null && (m[(16, 5)] = 16)
        and(line, and(s16, i6)) != null && (m[(16, 6)] = 16)
        and(line, and(s16, i7)) != null && (m[(16, 7)] = 16)
        and(line, and(s17, i0)) != null && (m[(17, 0)] = 17)
        and(line, and(s17, i1)) != null && (m[(17, 1)] = 17)
        and(line, and(s17, i2)) != null && (m[(17, 2)] = 18)
        and(line, and(s17, i3)) != null && (m[(17, 3)] = 17)
        and(line, and(s17, i4)) != null && (m[(17, 4)] = 17)
        and(line, and(s17, i5)) != null && (m[(17, 5)] = 17)
        and(line, and(s17, i6)) != null && (m[(17, 6)] = 17)
        and(line, and(s17, i7)) != null && (m[(17, 7)] = 17)
        and(line, and(s18, i0)) != null && (m[(18, 0)] = 18)
        and(line, and(s18, i1)) != null && (m[(18, 1)] = 18)
        and(line, and(s18, i2)) != null && (m[(18, 2)] = 19)
        and(line, and(s18, i3)) != null && (m[(18, 3)] = 18)
        and(line, and(s18, i4)) != null && (m[(18, 4)] = 18)
        and(line, and(s18, i5)) != null && (m[(18, 5)] = 18)
        and(line, and(s18, i6)) != null && (m[(18, 6)] = 7)
        and(line, and(s18, i7)) != null && (m[(18, 7)] = 18)
        and(line, and(s19, i0)) != null && (m[(19, 0)] = 19)
        and(line, and(s19, i1)) != null && (m[(19, 1)] = 19)
        and(line, and(s19, i2)) != null && (m[(19, 2)] = 19)
        and(line, and(s19, i3)) != null && (m[(19, 3)] = 19)
        and(line, and(s19, i4)) != null && (m[(19, 4)] = 19)
        and(line, and(s19, i5)) != null && (m[(19, 5)] = 23)
        and(line, and(s19, i6)) != null && (m[(19, 6)] = 19)
        and(line, and(s19, i7)) != null && (m[(19, 7)] = 20)
        and(line, and(s20, i0)) != null && (m[(20, 0)] = 20)
        and(line, and(s20, i1)) != null && (m[(20, 1)] = 20)
        and(line, and(s20, i2)) != null && (m[(20, 2)] = 20)
        and(line, and(s20, i3)) != null && (m[(20, 3)] = 21)
        and(line, and(s20, i4)) != null && (m[(20, 4)] = 20)
        and(line, and(s20, i5)) != null && (m[(20, 5)] = 20)
        and(line, and(s20, i6)) != null && (m[(20, 6)] = 21)
        and(line, and(s20, i7)) != null && (m[(20, 7)] = 20)
        and(line, and(s21, i0)) != null && (m[(21, 0)] = 22)
        and(line, and(s21, i1)) != null && (m[(21, 1)] = 21)
        and(line, and(s21, i2)) != null && (m[(21, 2)] = 21)
        and(line, and(s21, i3)) != null && (m[(21, 3)] = 21)
        and(line, and(s21, i4)) != null && (m[(21, 4)] = 21)
        and(line, and(s21, i5)) != null && (m[(21, 5)] = 21)
        and(line, and(s21, i6)) != null && (m[(21, 6)] = 21)
        and(line, and(s21, i7)) != null && (m[(21, 7)] = 21)
        and(line, and(s22, i0)) != null && (m[(22, 0)] = 22)
        and(line, and(s22, i1)) != null && (m[(22, 1)] = 22)
        and(line, and(s22, i2)) != null && (m[(22, 2)] = 22)
        and(line, and(s22, i3)) != null && (m[(22, 3)] = 22)
        and(line, and(s22, i4)) != null && (m[(22, 4)] = 22)
        and(line, and(s22, i5)) != null && (m[(22, 5)] = 22)
        and(line, and(s22, i6)) != null && (m[(22, 6)] = 22)
        and(line, and(s22, i7)) != null && (m[(22, 7)] = 23)
        and(line, and(s23, i0)) != null && (m[(23, 0)] = 23)
        and(line, and(s23, i1)) != null && (m[(23, 1)] = 29)
        and(line, and(s23, i2)) != null && (m[(23, 2)] = 23)
        and(line, and(s23, i3)) != null && (m[(23, 3)] = 23)
        and(line, and(s23, i4)) != null && (m[(23, 4)] = 23)
        and(line, and(s23, i5)) != null && (m[(23, 5)] = 24)
        and(line, and(s23, i6)) != null && (m[(23, 6)] = 23)
        and(line, and(s23, i7)) != null && (m[(23, 7)] = 23)
        and(line, and(s24, i0)) != null && (m[(24, 0)] = 24)
        and(line, and(s24, i1)) != null && (m[(24, 1)] = 25)
        and(line, and(s24, i2)) != null && (m[(24, 2)] = 24)
        and(line, and(s24, i3)) != null && (m[(24, 3)] = 24)
        and(line, and(s24, i4)) != null && (m[(24, 4)] = 24)
        and(line, and(s24, i5)) != null && (m[(24, 5)] = 24)
        and(line, and(s24, i6)) != null && (m[(24, 6)] = 24)
        and(line, and(s24, i7)) != null && (m[(24, 7)] = 9)
        and(line, and(s25, i0)) != null && (m[(25, 0)] = 25)
        and(line, and(s25, i1)) != null && (m[(25, 1)] = 25)
        and(line, and(s25, i2)) != null && (m[(25, 2)] = 26)
        and(line, and(s25, i3)) != null && (m[(25, 3)] = 25)
        and(line, and(s25, i4)) != null && (m[(25, 4)] = 25)
        and(line, and(s25, i5)) != null && (m[(25, 5)] = 25)
        and(line, and(s25, i6)) != null && (m[(25, 6)] = 25)
        and(line, and(s25, i7)) != null && (m[(25, 7)] = 25)
        and(line, and(s26, i0)) != null && (m[(26, 0)] = 26)
        and(line, and(s26, i1)) != null && (m[(26, 1)] = 26)
        and(line, and(s26, i2)) != null && (m[(26, 2)] = 27)
        and(line, and(s26, i3)) != null && (m[(26, 3)] = 26)
        and(line, and(s26, i4)) != null && (m[(26, 4)] = 26)
        and(line, and(s26, i5)) != null && (m[(26, 5)] = 26)
        and(line, and(s26, i6)) != null && (m[(26, 6)] = 26)
        and(line, and(s26, i7)) != null && (m[(26, 7)] = 26)
        and(line, and(s27, i0)) != null && (m[(27, 0)] = 27)
        and(line, and(s27, i1)) != null && (m[(27, 1)] = 27)
        and(line, and(s27, i2)) != null && (m[(27, 2)] = 27)
        and(line, and(s27, i3)) != null && (m[(27, 3)] = 27)
        and(line, and(s27, i4)) != null && (m[(27, 4)] = 27)
        and(line, and(s27, i5)) != null && (m[(27, 5)] = 27)
        and(line, and(s27, i6)) != null && (m[(27, 6)] = 27)
        and(line, and(s27, i7)) != null && (m[(27, 7)] = 28)
        and(line, and(s28, i0)) != null && (m[(28, 0)] = 28)
        and(line, and(s28, i1)) != null && (m[(28, 1)] = 28)
        and(line, and(s28, i2)) != null && (m[(28, 2)] = 28)
        and(line, and(s28, i3)) != null && (m[(28, 3)] = 29)
        and(line, and(s28, i4)) != null && (m[(28, 4)] = 28)
        and(line, and(s28, i5)) != null && (m[(28, 5)] = 28)
        and(line, and(s28, i6)) != null && (m[(28, 6)] = 29)
        and(line, and(s28, i7)) != null && (m[(28, 7)] = 28)
        and(line, and(s29, i0)) != null && (m[(29, 0)] = 30)
        and(line, and(s29, i1)) != null && (m[(29, 1)] = 29)
        and(line, and(s29, i2)) != null && (m[(29, 2)] = 29)
        and(line, and(s29, i3)) != null && (m[(29, 3)] = 29)
        and(line, and(s29, i4)) != null && (m[(29, 4)] = 29)
        and(line, and(s29, i5)) != null && (m[(29, 5)] = 29)
        and(line, and(s29, i6)) != null && (m[(29, 6)] = 29)
        and(line, and(s29, i7)) != null && (m[(29, 7)] = 29)
        and(line, and(s30, i0)) != null && (m[(30, 0)] = 30)
        and(line, and(s30, i1)) != null && (m[(30, 1)] = 30)
        and(line, and(s30, i2)) != null && (m[(30, 2)] = 30)
        and(line, and(s30, i3)) != null && (m[(30, 3)] = 30)
        and(line, and(s30, i4)) != null && (m[(30, 4)] = 30)
        and(line, and(s30, i5)) != null && (m[(30, 5)] = 30)
        and(line, and(s30, i6)) != null && (m[(30, 6)] = 30)
        and(line, and(s30, i7)) != null && (m[(30, 7)] = 31)
        and(line, and(s31, i0)) != null && (m[(31, 0)] = 31)
        and(line, and(s31, i1)) != null && (m[(31, 1)] = 31)
        and(line, and(s31, i2)) != null && (m[(31, 2)] = 0)
        and(line, and(s31, i3)) != null && (m[(31, 3)] = 31)
        and(line, and(s31, i4)) != null && (m[(31, 4)] = 31)
        and(line, and(s31, i5)) != null && (m[(31, 5)] = 0)
        and(line, and(s31, i6)) != null && (m[(31, 6)] = 31)
        and(line, and(s31, i7)) != null && (m[(31, 7)] = 31)
        
        return m
        
    end

    decomp_ns16 = decomp(ns16_s)
    decomp_ns8  = decomp(ns8_s)
    decomp_ns4  = decomp(ns4_s)
    decomp_ns2  = decomp(ns2_s)
    decomp_ns1  = decomp(ns1_s)

    return (decomp_ns16, decomp_ns8, decomp_ns4, decomp_ns2, decomp_ns1)
    
end # of ns2dict()


# dict2ps ==============================================================
# ======================================================================

function dict2ps(decomp_ns16, decomp_ns8, decomp_ns4, decomp_ns2,
                 decomp_ns1)
    
    # The following if-elif ladder illustrates a technique of
    # moving foward to a next time-frame by passing fault as a
    # fault-pattern (U.S. Patent No. 8,156,395).
    
    # initialize the ps_ variables
    ps_16 = 0
    ps_8 = 0
    ps_4 = 0
    ps_2 = 0
    ps_1 = 0

    # define a function for computing a fault_pattern (fp)
    function fault_pattern(ps_16, ps_8, ps_4, ps_2, ps_1)
        fp = 0
        ps_16 != 33 && (fp += 16)
        ps_8 != 33 && (fp += 8)
        ps_4 != 33 && (fp += 4)
        ps_2 != 33 && (fp += 2)
        ps_1 != 33 && (fp += 1)
        return fp
    end
    
    empty = Dict()

    if decomp_ns16 != empty
        a = pop!(decomp_ns16)               # a = ((s, i), ns)
        a_key = a[1]                        # a_key = (s, i)
        ps_16 = a[2]                        # ps_16 = ns
        ps_8 = pop!(decomp_ns8, a_key, 33)
        ps_4 = pop!(decomp_ns4, a_key, 33)
        ps_2 = pop!(decomp_ns2, a_key, 33)
        ps_1 = pop!(decomp_ns1, a_key, 33)
        fp = fault_pattern(ps_16, ps_8, ps_4, ps_2, ps_1)
        chosen_item = (fp, a)

    elseif decomp_ns8 != empty
        ps_16 = 33
        a = pop!(decomp_ns8)
        a_key = a[1]
        ps_8 = a[2]
        ps_4 = pop!(decomp_ns4, a_key, 33)
        ps_2 = pop!(decomp_ns2, a_key, 33)
        ps_1 = pop!(decomp_ns1, a_key, 33)
        fp = fault_pattern(ps_16, ps_8, ps_4, ps_2, ps_1)
        chosen_item = (fp, a)

    elseif decomp_ns4 != empty
        ps_16 = 33
        ps_8 = 33
        a = pop!(decomp_ns4)
        a_key = a[1]
        ps_4 = a[2]
        ps_2 = pop!(decomp_ns2, a_key, 33)
        ps_1 = pop!(decomp_ns1, a_key, 33)
        fp = fault_pattern(ps_16, ps_8, ps_4, ps_2, ps_1)
        chosen_item = (fp, a)

    elseif decomp_ns2 != empty
        ps_16 = 33
        ps_8 = 33
        ps_4 = 33
        a = pop!(decomp_ns2)
        a_key = a[1]
        ps_2 = a[2]
        ps_1 = pop!(decomp_ns1, a_key, 33)
        fp = fault_pattern(ps_16, ps_8, ps_4, ps_2, ps_1)
        chosen_item = (fp, a)

    elseif decomp_ns1 != empty
        ps_16 = 33
        ps_8 = 33
        ps_4 = 33
        ps_2 = 33
        a = pop!(decomp_ns1)
        ps_1 = a[2]
        fp = fault_pattern(ps_16, ps_8, ps_4, ps_2, ps_1)
        chosen_item = (fp, a)

    else
        ps_16 = 33
        ps_8 = 33
        ps_4 = 33
        ps_2 = 33
        ps_1 = 33
        chosen_item = ()

    end
    
    (ps16_s, ps8_s, ps4_s, ps2_s, ps1_s) =
        setup_PS_lines(ps_16, ps_8, ps_4, ps_2, ps_1)


    return (ps16_s, ps8_s, ps4_s, ps2_s, ps1_s, chosen_item,
            decomp_ns16, decomp_ns8, decomp_ns4, decomp_ns2, decomp_ns1)

end # of dict2ps()


# COMMON SETUP PS LINES ================================================
# ======================================================================

function setup_PS_lines(g16, g8, g4, g2, g1)

    g16 == 33 && (h16 = no_terms)
    g16 == 0 && (h16 = s0)
    g16 == 1 && (h16 = s1)
    g16 == 2 && (h16 = s2)
    g16 == 3 && (h16 = s3)
    g16 == 4 && (h16 = s4)
    g16 == 5 && (h16 = s5)
    g16 == 6 && (h16 = s6)
    g16 == 7 && (h16 = s7)
    g16 == 8 && (h16 = s8)
    g16 == 9 && (h16 = s9)
    g16 == 10 && (h16 = s10)
    g16 == 11 && (h16 = s11)
    g16 == 12 && (h16 = s12)
    g16 == 13 && (h16 = s13)
    g16 == 14 && (h16 = s14)
    g16 == 15 && (h16 = s15)
    g16 == 16 && (h16 = s16)
    g16 == 17 && (h16 = s17)
    g16 == 18 && (h16 = s18)
    g16 == 19 && (h16 = s19)
    g16 == 20 && (h16 = s20)
    g16 == 21 && (h16 = s21)
    g16 == 22 && (h16 = s22)
    g16 == 23 && (h16 = s23)
    g16 == 24 && (h16 = s24)
    g16 == 25 && (h16 = s25)
    g16 == 26 && (h16 = s26)
    g16 == 27 && (h16 = s27)
    g16 == 28 && (h16 = s28)
    g16 == 29 && (h16 = s29)
    g16 == 30 && (h16 = s30)
    g16 == 31 && (h16 = s31)

    g8 == 33 && (h8 = no_terms)
    g8 == 0 && (h8 = s0)
    g8 == 1 && (h8 = s1)
    g8 == 2 && (h8 = s2)
    g8 == 3 && (h8 = s3)
    g8 == 4 && (h8 = s4)
    g8 == 5 && (h8 = s5)
    g8 == 6 && (h8 = s6)
    g8 == 7 && (h8 = s7)
    g8 == 8 && (h8 = s8)
    g8 == 9 && (h8 = s9)
    g8 == 10 && (h8 = s10)
    g8 == 11 && (h8 = s11)
    g8 == 12 && (h8 = s12)
    g8 == 13 && (h8 = s13)
    g8 == 14 && (h8 = s14)
    g8 == 15 && (h8 = s15)
    g8 == 16 && (h8 = s16)
    g8 == 17 && (h8 = s17)
    g8 == 18 && (h8 = s18)
    g8 == 19 && (h8 = s19)
    g8 == 20 && (h8 = s20)
    g8 == 21 && (h8 = s21)
    g8 == 22 && (h8 = s22)
    g8 == 23 && (h8 = s23)
    g8 == 24 && (h8 = s24)
    g8 == 25 && (h8 = s25)
    g8 == 26 && (h8 = s26)
    g8 == 27 && (h8 = s27)
    g8 == 28 && (h8 = s28)
    g8 == 29 && (h8 = s29)
    g8 == 30 && (h8 = s30)
    g8 == 31 && (h8 = s31)

    g4 == 33 && (h4 = no_terms)
    g4 == 0 && (h4 = s0)
    g4 == 1 && (h4 = s1)
    g4 == 2 && (h4 = s2)
    g4 == 3 && (h4 = s3)
    g4 == 4 && (h4 = s4)
    g4 == 5 && (h4 = s5)
    g4 == 6 && (h4 = s6)
    g4 == 7 && (h4 = s7)
    g4 == 8 && (h4 = s8)
    g4 == 9 && (h4 = s9)
    g4 == 10 && (h4 = s10)
    g4 == 11 && (h4 = s11)
    g4 == 12 && (h4 = s12)
    g4 == 13 && (h4 = s13)
    g4 == 14 && (h4 = s14)
    g4 == 15 && (h4 = s15)
    g4 == 16 && (h4 = s16)
    g4 == 17 && (h4 = s17)
    g4 == 18 && (h4 = s18)
    g4 == 19 && (h4 = s19)
    g4 == 20 && (h4 = s20)
    g4 == 21 && (h4 = s21)
    g4 == 22 && (h4 = s22)
    g4 == 23 && (h4 = s23)
    g4 == 24 && (h4 = s24)
    g4 == 25 && (h4 = s25)
    g4 == 26 && (h4 = s26)
    g4 == 27 && (h4 = s27)
    g4 == 28 && (h4 = s28)
    g4 == 29 && (h4 = s29)
    g4 == 30 && (h4 = s30)
    g4 == 31 && (h4 = s31)

    g2 == 33 && (h2 = no_terms)
    g2 == 0 && (h2 = s0)
    g2 == 1 && (h2 = s1)
    g2 == 2 && (h2 = s2)
    g2 == 3 && (h2 = s3)
    g2 == 4 && (h2 = s4)
    g2 == 5 && (h2 = s5)
    g2 == 6 && (h2 = s6)
    g2 == 7 && (h2 = s7)
    g2 == 8 && (h2 = s8)
    g2 == 9 && (h2 = s9)
    g2 == 10 && (h2 = s10)
    g2 == 11 && (h2 = s11)
    g2 == 12 && (h2 = s12)
    g2 == 13 && (h2 = s13)
    g2 == 14 && (h2 = s14)
    g2 == 15 && (h2 = s15)
    g2 == 16 && (h2 = s16)
    g2 == 17 && (h2 = s17)
    g2 == 18 && (h2 = s18)
    g2 == 19 && (h2 = s19)
    g2 == 20 && (h2 = s20)
    g2 == 21 && (h2 = s21)
    g2 == 22 && (h2 = s22)
    g2 == 23 && (h2 = s23)
    g2 == 24 && (h2 = s24)
    g2 == 25 && (h2 = s25)
    g2 == 26 && (h2 = s26)
    g2 == 27 && (h2 = s27)
    g2 == 28 && (h2 = s28)
    g2 == 29 && (h2 = s29)
    g2 == 30 && (h2 = s30)
    g2 == 31 && (h2 = s31)

    g1 == 33 && (h1 = no_terms)
    g1 == 0 && (h1 = s0)
    g1 == 1 && (h1 = s1)
    g1 == 2 && (h1 = s2)
    g1 == 3 && (h1 = s3)
    g1 == 4 && (h1 = s4)
    g1 == 5 && (h1 = s5)
    g1 == 6 && (h1 = s6)
    g1 == 7 && (h1 = s7)
    g1 == 8 && (h1 = s8)
    g1 == 9 && (h1 = s9)
    g1 == 10 && (h1 = s10)
    g1 == 11 && (h1 = s11)
    g1 == 12 && (h1 = s12)
    g1 == 13 && (h1 = s13)
    g1 == 14 && (h1 = s14)
    g1 == 15 && (h1 = s15)
    g1 == 16 && (h1 = s16)
    g1 == 17 && (h1 = s17)
    g1 == 18 && (h1 = s18)
    g1 == 19 && (h1 = s19)
    g1 == 20 && (h1 = s20)
    g1 == 21 && (h1 = s21)
    g1 == 22 && (h1 = s22)
    g1 == 23 && (h1 = s23)
    g1 == 24 && (h1 = s24)
    g1 == 25 && (h1 = s25)
    g1 == 26 && (h1 = s26)
    g1 == 27 && (h1 = s27)
    g1 == 28 && (h1 = s28)
    g1 == 29 && (h1 = s29)
    g1 == 30 && (h1 = s30)
    g1 == 31 && (h1 = s31)

    return (h16, h8, h4, h2, h1)
    
end # of setup_ps_lines()


# peek =================================================================
# ======================================================================

function peek(fp::Int64, s::Int64)
    
    # How to use peek. Choose a fault. Execute peek(0,0). If you have
    # an output, read the S/I. That's your test. Else, select a tile
    # from the list and execute peek using the (fp, ns) for your
    # chosen tile. If peek returns an output, use the sequence of S/Is
    # from the tiles you have selected. That's your test. Else, choose
    # a tile and repeat.
    
    sx = convert_state(s)
    
    (pt16, pt8, pt4, pt2, pt1) = convert_key(fp, sx)
    
    println("cudd_Large,  ", fault)  #, "/", fault_B)
    
    ##############################################################
    (out4_s, out2_s, out1_s, ns16_s, ns8_s, ns4_s, ns2_s, ns1_s) =
        ps2ns(pt16, pt8, pt4, pt2, pt1, fault)
    ##############################################################
    
    out4_tmp = allSAT(out4_s)
    println("out4_s:", out4_tmp)
    out2_tmp = allSAT(out2_s)
    println("out2_s:", out4_tmp)
    out1_tmp = allSAT(out1_s)
    println("out1_s:", out1_tmp)
    
    ###############################################################
    (decomp_ns16, decomp_ns8, decomp_ns4, decomp_ns2, decomp_ns1) =
        ns2dict(ns16_s, ns8_s, ns4_s, ns2_s, ns1_s)
    ###############################################################
    
    empty = Dict()

    n = []
    
    while ((decomp_ns16 != empty) || (decomp_ns8 != empty) ||
        (decomp_ns4 != empty) || (decomp_ns2 != empty) ||
        (decomp_ns1 != empty))
        
        ################################################################
        (ps16_s, ps8_s, ps4_s, ps2_s, ps1_s, chosen_item,decomp_ns16,
            decomp_ns8, decomp_ns4, decomp_ns2, decomp_ns1) =
            dict2ps(decomp_ns16, decomp_ns8, decomp_ns4,
            decomp_ns2, decomp_ns1)
        ################################################################
        
        # chosen_item = (fpout, ((sin, i), ns))
        a = chosen_item
        # tile format: link tile: (fp, s) a[2][1] (a[1], a[2][2])
        #     "      output tile: (fp, s) {output_no., s/i}, a[2][2]
        
        
        fp_fp = a[1]
        s_fp = a[2][1][1]
        i_fp = a[2][1][2]
        ns_fp = a[2][2]

        
        # ================================================================================================
        
        if (fp == 0) || (s_fp == s)  # first time-frame, consistent state
            println("(",fp,", ",s,")           \"s",s_fp,"i",i_fp,"\"   (",fp_fp,", ",ns_fp,")")
            push!(n,(s_fp, i_fp, fp_fp, ns_fp))
        end
        
        # ================================================================================================

    end
    
    return n

end # 0f peek()


# helpers ==============================================================
# ======================================================================

function convert_state(s::Int64)
    
    if s == 0
        sx = s0
    elseif s == 1
        sx = s1
    elseif s == 2
        sx = s2
    elseif s == 3
        sx = s3
    elseif s == 4
        sx = s4
    elseif s == 5
        sx = s5
    elseif s == 6
        sx = s6
    elseif s == 7
        sx = s7
    elseif s == 8
        sx = s8
    elseif s == 9
        sx = s9
    elseif s == 10
        sx = s10
    elseif s == 11
        sx = s11
    elseif s == 12
        sx = s12
    elseif s == 13
        sx = s13
    elseif s == 14
        sx = s14
    elseif s == 15
        sx = s15
    elseif s == 16
        sx = s16
    elseif s == 17
        sx = s17
    elseif s == 18
        sx = s18
    elseif s == 19
        sx = s19
    elseif s == 20
        sx = s20
    elseif s == 21
        sx = s21
    elseif s == 22
        sx = s22
    elseif s == 23
        sx = s23
    elseif s == 24
        sx = s24
    elseif s == 25
        sx = s25
    elseif s == 26
        sx = s26
    elseif s == 27
        sx = s27
    elseif s == 28
        sx = s28
    elseif s == 29
        sx = s29
    elseif s == 30
        sx = s30
    else
        sx = s31
    end
    
    return sx
end # of convert_state()


function convert_key(fp::Int64, sx::Ptr{Nothing})
    
    b = ()
    
    if fp == 0
        b = (null, null, null, null, null)
    elseif fp == 1
        b = (null, null, null, null, sx)
    elseif fp == 2
        b = (null, null, null, sx, null)
    elseif fp == 3
        b = (null, null, null, sx, sx)
    elseif fp == 4
        b = (null, null, sx, null, null)
    elseif fp == 5
        b = (null, null, sx, null, sx)
    elseif fp == 6
        b = (null, null, sx, sx, null)
    elseif fp == 7
        b = (null, null, sx, sx, sx)
    elseif fp == 8
        b = (null, sx, null, null, null)
    elseif fp == 9
        b = (null, sx, null, null, sx)
    elseif fp == 10
        b = (null, sx, null, sx, null)
    elseif fp == 11
        b = (null, sx, null, sx, sx)
    elseif fp == 12
        b = (null, sx, sx, null, null)
    elseif fp == 13
        b = (null, sx, sx, null, sx)
    elseif fp == 14
        b = (null, sx, sx, sx, null)
    elseif fp == 15
        b = (null, sx, sx, sx, sx)
    elseif fp == 16
        b = (sx, null, null, null, null)
    elseif fp == 17
        b = (sx, null, null, null, sx)
    elseif fp == 18
        b = (sx, null, null, sx, null)
    elseif fp == 19
        b = (sx, null, null, sx, sx)
    elseif fp == 20
        b = (sx, null, sx, null, null)
    elseif fp == 21
        b = (sx, null, sx, null, sx)
    elseif fp == 22
        b = (sx, null, sx, sx, null)
    elseif fp == 23
        b = (sx, null, sx, sx, sx)
    elseif fp == 24
        b = (sx, sx, null, null, null)
    elseif fp == 25
        b = (sx, sx, null, null, sx)
    elseif fp == 26
        b = (sx, sx, null, sx, null)
    elseif fp == 27
        b = (sx, sx, null, sx, sx)
    elseif fp == 28
        b = (sx, sx, sx, null, null)
    elseif fp == 29
        b = (sx, sx, sx, null, sx)
    elseif fp == 30
        b = (sx, sx, sx, sx, null)
    else
        b = (sx, sx, sx, sx, sx)
    end
    
    return b
end # of convert_key()


function si2str(s::Int64, i::Int64)
    "s" * string(s, base = 10) * "i" * string(i, base = 10)
end # of si2str()

#=
# ON_set ===============================================================
# ======================================================================

function ON_set(ps16_1, ps8_1, ps4_1, ps2_1, ps1_1,
                in4_1, in2_1, in1_1, fault_B)

    # level 0

    fault_B == "ps16:0" && (ps16_1 = ZERO)
    fault_B == "ps16:1" && (ps16_1 = ONE)

    fault_B == "ps8:0" && (ps8_1 = ZERO)
    fault_B == "ps8:1" && (ps8_1 = ONE)

    fault_B == "ps4:0" && (ps4_1 = ZERO)
    fault_B == "ps4:1" && (ps4_1 = ONE)

    fault_B == "ps2:0" && (ps2_1 = ZERO)
    fault_B == "ps2:1" && (ps2_1 = ONE)

    fault_B == "ps1:0" && (ps1_1 = ZERO)
    fault_B == "ps1:1" && (ps1_1 = ONE)

    fault_B == "in4:0" && (in4_1 = ZERO)
    fault_B == "in4:1" && (in4_1 = ONE)

    fault_B == "in2:0" && (in2_1 = ZERO)
    fault_B == "in2:1" && (in2_1 = ONE)

    fault_B == "in1:0" && (in1_1 = ZERO)
    fault_B == "in1:1" && (in1_1 = ONE)
    
    # level 1

    _,nps1_1 = op1not(null, ps1_1)
    fault_B == "nps1:0" && (nps1_1 = ZERO)
    fault_B == "nps1:1" && (nps1_1 = ONE)

    _,nps2_1 = op1not(null, ps2_1)
    fault_B == "nps2:0" && (nps2_1 = ZERO)
    fault_B == "nps2:1" && (nps2_1 = ONE)

    _,nps4_1 = op1not(null, ps4_1)
    fault_B == "nps4:0" && (nps4_1 = ZERO)
    fault_B == "nps4:1" && (nps4_1 = ONE)

    _,nps8_1 = op1not(null, ps8_1)
    fault_B == "nps8:0" && (nps8_1 = ZERO)
    fault_B == "nps8:1" && (nps8_1 = ONE)

    _,nps16_1 = op1not(null,ps16_1)
    fault_B == "nps16:0" && (nps16_1 = ZERO)
    fault_B == "nps16:1" && (nps16_1 = ONE)

    _,nin1_1 = op1not(null, in1_1)
    fault_B == "nin1:0" && (nin1_1 = ZERO)
    fault_B == "nin1:1" && (nin1_1 = ONE)

    _,nin2_1 = op1not(null, in2_1)
    fault_B == "nin2:0" && (nin2_1 = ZERO)
    fault_B == "nin2:1" && (nin2_1 = ONE)

    _,nin4_1 = op1not(null, in4_1)
    fault_B == "nin4:0" && (nin4_1 = ZERO)
    fault_B == "nin4:1" && (nin4_1 = ONE)

    _,i7_1 = op3and(null,null,null,in4_1,in2_1,in1_1)
    fault_B == "i7:0" && (i7_1 = ZERO)
    fault_B == "i7:1" && (i7_1 = ONE)
    
    # level 2

    _,i0_1 = op3and(null,null,null,nin4_1,nin2_1,nin1_1)
    fault_B == "i0:0" && (i0_1 = ZERO)
    fault_B == "i0:1" && (i0_1 = ONE)

    _,i1_1 = op3and(null,null,null,nin4_1,nin2_1,in1_1)
    fault_B == "i1:0" && (i1_1 = ZERO)
    fault_B == "i1:1" && (i1_1 = ONE)

    _,i2_1 = op3and(null,null,null,nin4_1,in2_1,nin1_1)
    fault_B == "i2:0" && (i2_1 = ZERO)
    fault_B == "i2:1" && (i2_1 = ONE)

    _,i3_1 = op3and(null,null,null,nin4_1,in2_1,in1_1)
    fault_B == "i3:0" && (i3_1 = ZERO)
    fault_B == "i3:1" && (i3_1 = ONE)

    _,i4_1 = op3and(null,null,null,in4_1,nin2_1,nin1_1)
    fault_B == "i4:0" && (i4_1 = ZERO)
    fault_B == "i4:1" && (i4_1 = ONE)

    _,i5_1 = op3and(null,null,null,in4_1,nin2_1,in1_1)
    fault_B == "i5:0" && (i5_1 = ZERO)
    fault_B == "i5:1" && (i5_1 = ONE)

    _,i6_1 = op3and(null,null,null,in4_1,in2_1,nin1_1)
    fault_B == "i6:0" && (i6_1 = ZERO)
    fault_B == "i6:1" && (i6_1 = ONE)

    _,ls0_1 = op3and(null,null,null,nps4_1,nps2_1,nps1_1)
    fault_B == "ls0:0" && (ls0_1 = ZERO)
    fault_B == "ls0:1" && (ls0_1 = ONE)

    _,ls1_1 = op3and(null,null,null,nps4_1,nps2_1,ps1_1)
    fault_B == "ls1:0" && (ls1_1 = ZERO)
    fault_B == "ls1:1" && (ls1_1 = ONE)

    _,ls2_1 = op3and(null,null,null,nps4_1,ps2_1,nps1_1)
    fault_B == "ls2:0" && (ls2_1 = ZERO)
    fault_B == "ls2:1" && (ls2_1 = ONE)

    _,ls3_1 = op3and(null,null,null,nps4_1,ps2_1,ps1_1)
    fault_B == "ls3:0" && (ls3_1 = ZERO)
    fault_B == "ls3:1" && (ls3_1 = ONE)

    _,ls4_1 = op3and(null,null,null,ps4_1,nps2_1,nps1_1)
    fault_B == "ls4:0" && (ls4_1 = ZERO)
    fault_B == "ls4:1" && (ls4_1 = ONE)

    _,ls5_1 = op3and(null,null,null,ps4_1,nps2_1,ps1_1)
    fault_B == "ls5:0" && (ls5_1 = ZERO)
    fault_B == "ls5:1" && (ls5_1 = ONE)

    _,ls6_1 = op3and(null,null,null,ps4_1,ps2_1,nps1_1)
    fault_B == "ls6:0" && (ls6_1 = ZERO)
    fault_B == "ls6:1" && (ls6_1 = ONE)

    _,ls7_1 = op3and(null,null,null,ps4_1,ps2_1,ps1_1)
    fault_B == "ls7:0" && (ls7_1 = ZERO)
    fault_B == "ls7:1" && (ls7_1 = ONE)

    _,ni7_1 = op1not(null,i7_1)
    fault_B == "ni7:0" && (ni7_1 = ZERO)
    fault_B == "ni7:1" && (ni7_1 = ONE)

    _,s31_1 = op3and(null,null,null,ps16_1,ps8_1,ls7_1)
    fault_B == "s31:0" && (s31_1 = ZERO)
    fault_B == "s31:1" && (s31_1 = ONE)
    
    # level 3

    _,ni0_1 = op1not(null,i0_1)
    fault_B == "ni0:0" && (ni0_1 = ZERO)
    fault_B == "ni0:1" && (ni0_1 = ONE)

    _,ni1_1 = op1not(null,i1_1)
    fault_B == "ni1:0" && (ni1_1 = ZERO)
    fault_B == "ni1:1" && (ni1_1 = ONE)

    _,ni2_1 = op1not(null,i2_1)
    fault_B == "ni2:0" && (ni2_1 = ZERO)
    fault_B == "ni2:1" && (ni2_1 = ONE)

    _,ni3_1 = op1not(null,i3_1)
    fault_B == "ni3:0" && (ni3_1 = ZERO)
    fault_B == "ni3:1" && (ni3_1 = ONE)

    _,ni5_1 = op1not(null,i5_1)
    fault_B == "ni5:0" && (ni5_1 = ZERO)
    fault_B == "ni5:1" && (ni5_1 = ONE)

    _,ni6_1 = op1not(null,i6_1)
    fault_B == "ni6:0" && (ni6_1 = ZERO)
    fault_B == "ni6:1" && (ni6_1 = ONE)

    _,s0_1 = op3and(null,null,null,nps16_1,nps8_1,ls0_1)
    fault_B == "s0:0" && (s0_1 = ZERO)
    fault_B == "s0:1" && (s0_1 = ONE)

    _,s1_1 = op3and(null,null,null,nps16_1,nps8_1,ls1_1)
    fault_B == "s1:0" && (s1_1 = ZERO)
    fault_B == "s1:1" && (s1_1 = ONE)

    _,s2_1 = op3and(null,null,null,nps16_1,nps8_1,ls2_1)
    fault_B == "s2:0" && (s2_1 = ZERO)
    fault_B == "s2:1" && (s2_1 = ONE)

    _,s3_1 = op3and(null,null,null,nps16_1,nps8_1,ls3_1)
    fault_B == "s3:0" && (s3_1 = ZERO)
    fault_B == "s3:1" && (s3_1 = ONE)

    _,s4_1 = op3and(null,null,null,nps16_1,nps8_1,ls4_1)
    fault_B == "s4:0" && (s4_1 = ZERO)
    fault_B == "s4:1" && (s4_1 = ONE)

    _,s5_1 = op3and(null,null,null,nps16_1,nps8_1,ls5_1)
    fault_B == "s5:0" && (s5_1 = ZERO)
    fault_B == "s5:1" && (s5_1 = ONE)

    _,s6_1 = op3and(null,null,null,nps16_1,nps8_1,ls6_1)
    fault_B == "ls7:0" && (ls7_1 = ZERO)
    fault_B == "ls7:1" && (ls7_1 = ONE)

    _,s7_1 = op3and(null,null,null,nps16_1,nps8_1,ls7_1)
    fault_B == "s7:0" && (s7_1 = ZERO)
    fault_B == "s7:1" && (s7_1 = ONE)

    _,s8_1 = op3and(null,null,null,nps16_1,ps8_1,ls0_1)
    fault_B == "s8:0" && (s8_1 = ZERO)
    fault_B == "s8:1" && (s8_1 = ONE)

    _,s9_1 = op3and(null,null,null,nps16_1,ps8_1,ls1_1)
    fault_B == "s9:0" && (s9_1 = ZERO)
    fault_B == "s9:1" && (s9_1 = ONE)

    _,s10_1 = op3and(null,null,null,nps16_1,ps8_1,ls2_1)
    fault_B == "s10:0" && (s10_1 = ZERO)
    fault_B == "s10:1" && (s10_1 = ONE)

    _,s11_1 = op3and(null,null,null,nps16_1,ps8_1,ls3_1)
    fault_B == "s11:0" && (s11_1 = ZERO)
    fault_B == "s11:1" && (s11_1 = ONE)

    _,s12_1 = op3and(null,null,null,nps16_1,ps8_1,ls4_1)
    fault_B == "s12:0" && (s12_1 = ZERO)
    fault_B == "s12:1" && (s12_1 = ONE)

    _,s13_1 = op3and(null,null,null,nps16_1,ps8_1,ls5_1)
    fault_B == "s13:0" && (s13_1 = ZERO)
    fault_B == "s13:1" && (s13_1 = ONE)

    _,s14_1 = op3and(null,null,null,nps16_1,ps8_1,ls6_1)
    fault_B == "s14:0" && (s14_1 = ZERO)
    fault_B == "s14:1" && (s14_1 = ONE)

    _,s15_1 = op3and(null,null,null,nps16_1,ps8_1,ls7_1)
    fault_B == "s15:0" && (s15_1 = ZERO)
    fault_B == "s15:1" && (s15_1 = ONE)

    _,s16_1 = op3and(null,null,null,ps16_1,nps8_1,ls0_1)
    fault_B == "s16:0" && (s16_1 = ZERO)
    fault_B == "s16:1" && (s16_1 = ONE)

    _,s17_1 = op3and(null,null,null,ps16_1,nps8_1,ls1_1)
    fault_B == "s17:0" && (s17_1 = ZERO)
    fault_B == "s17:1" && (s17_1 = ONE)

    _,s18_1 = op3and(null,null,null,ps16_1,nps8_1,ls2_1)
    fault_B == "s18:0" && (s18_1 = ZERO)
    fault_B == "s18:1" && (s18_1 = ONE)

    _,s19_1 = op3and(null,null,null,ps16_1,nps8_1,ls3_1)
    fault_B == "s19:0" && (s19_1 = ZERO)
    fault_B == "s19:1" && (s19_1 = ONE)

    _,s20_1 = op3and(null,null,null,ps16_1,nps8_1,ls4_1)
    fault_B == "s20:0" && (s20_1 = ZERO)
    fault_B == "s20:1" && (s20_1 = ONE)

    _,s21_1 = op3and(null,null,null,ps16_1,nps8_1,ls5_1)
    fault_B == "s21:0" && (s21_1 = ZERO)
    fault_B == "s21:1" && (s21_1 = ONE)

    _,s22_1 = op3and(null,null,null,ps16_1,nps8_1,ls6_1)
    fault_B == "s22:0" && (s22_1 = ZERO)
    fault_B == "s22:1" && (s22_1 = ONE)

    _,s23_1 = op3and(null,null,null,ps16_1,nps8_1,ls7_1)
    fault_B == "s23:0" && (s23_1 = ZERO)
    fault_B == "s23:1" && (s23_1 = ONE)

    _,s24_1 = op3and(null,null,null,ps16_1,ps8_1,ls0_1)
    fault_B == "s24:0" && (s24_1 = ZERO)
    fault_B == "s24:1" && (s24_1 = ONE)

    _,s25_1 = op3and(null,null,null,ps16_1,ps8_1,ls1_1)
    fault_B == "s25:0" && (s25_1 = ZERO)
    fault_B == "s25:1" && (s25_1 = ONE)

    _,s26_1 = op3and(null,null,null,ps16_1,ps8_1,ls2_1)
    fault_B == "s26:0" && (s26_1 = ZERO)
    fault_B == "s26:1" && (s26_1 = ONE)

    _,s27_1 = op3and(null,null,null,ps16_1,ps8_1,ls3_1)
    fault_B == "s27:0" && (s27_1 = ZERO)
    fault_B == "s27:1" && (s27_1 = ONE)

    _,s28_1 = op3and(null,null,null,ps16_1,ps8_1,ls4_1)
    fault_B == "s28:0" && (s28_1 = ZERO)
    fault_B == "s28:1" && (s28_1 = ONE)

    _,s29_1 = op3and(null,null,null,ps16_1,ps8_1,ls5_1)
    fault_B == "s29:0" && (s29_1 = ZERO)
    fault_B == "s29:1" && (s29_1 = ONE)

    _,s30_1 = op3and(null,null,null,ps16_1,ps8_1,ls6_1)
    fault_B == "s30:0" && (s30_1 = ZERO)
    fault_B == "s30:1" && (s30_1 = ONE)

    _,b2_1 = op3or(null,null,null,i5_1,i3_1,i2_1)
    fault_B == "b2:0" && (b2_1 = ZERO)
    fault_B == "b2:1" && (b2_1 = ONE)

    _,b7_1 = op2or(null,null,i5_1,i1_1)
    fault_B == "b7:0" && (b7_1 = ZERO)
    fault_B == "b7:1" && (b7_1 = ONE)

    _,c5_1 = op2or(null,null,i7_1,i5_1)
    fault_B == "c5:0" && (c5_1 = ZERO)
    fault_B == "c5:1" && (c5_1 = ONE)

    _,c13_1 = op2or(null,null,i3_1,i2_1)
    fault_B == "c13:0" && (c13_1 = ZERO)
    fault_B == "c13:1" && (c13_1 = ONE)

    _,e5_1 = op2or(null,null,i5_1,i4_1)
    fault_B == "e5:0" && (e5_1 = ZERO)
    fault_B == "e5:1" && (e5_1 = ONE)

    _,e10_1 = op3or(null,null,null,i6_1,i2_1,i0_1)
    fault_B == "e10:0" && (e10_1 = ZERO)
    fault_B == "e10:1" && (e10_1 = ONE)

    _,e12_1 = op2or(null,null,i6_1,i3_1)
    fault_B == "e12:0" && (e12_1 = ZERO)
    fault_B == "e12:1" && (e12_1 = ONE)

    _,e18_1 = op2or(null,null,i6_1,i2_1)
    fault_B == "e18:0" && (e18_1 = ZERO)
    fault_B == "e18:1" && (e18_1 = ONE)

    _,e24_1 = op2or(null,null,i7_1,i1_1)
    fault_B == "e24:0" && (e24_1 = ZERO)
    fault_B == "e24:1" && (e24_1 = ONE)
    
    # level 4

    _,a1_1 = op2and(null,null,s10_1,i0_1)
    fault_B == "a1:0" && (a1_1 = ZERO)
    fault_B == "a1:1" && (a1_1 = ONE)

    _,a2_1 = op2and(null,null,s15_1,i5_1)
    fault_B == "a2:0" && (a2_1 = ZERO)
    fault_B == "a2:1" && (a2_1 = ONE)

    _,a3_1 = op2and(null,null,s18_1,ni6_1)
    fault_B == "a3:0" && (a3_1 = ZERO)
    fault_B == "a3:1" && (a3_1 = ONE)

    _,a4_1 = op3or(null,null,null,s20_1,s21_1,s22_1)
    fault_B == "a4:0" && (a4_1 = ZERO)
    fault_B == "a4:1" && (a4_1 = ONE)

    _,a5_1 = op2and(null,null,s24_1,ni7_1)
    fault_B == "a5:0" && (a5_1 = ZERO)
    fault_B == "a5:1" && (a5_1 = ONE)

    _,a6_1 = op2or(null,null,s25_1,s26_1)
    fault_B == "a6:0" && (a6_1 = ZERO)
    fault_B == "a6:1" && (a6_1 = ONE)

    _,a7_1 = op3or(null,null,null,s27_1,s28_1,s29_1)
    fault_B == "a7:0" && (a7_1 = ZERO)
    fault_B == "a7:1" && (a7_1 = ONE)

    _,a9_1 = op3and(null,null,null,s31_1,ni5_1,ni2_1)
    fault_B == "a9:0" && (a9_1 = ZERO)
    fault_B == "a9:1" && (a9_1 = ONE)

    _,b1_1 = op2and(null,null,s3_1,i2_1)
    fault_B == "b1:0" && (b1_1 = ZERO)
    fault_B == "b1:1" && (b1_1 = ONE)

    _,b3_1 = op2and(null,null,b2_1,s7_1)
    fault_B == "b3:0" && (b3_1 = ZERO)
    fault_B == "b3:1" && (b3_1 = ONE)

    _,b4_1 = op2and(null,null,s10_1,ni6_1)
    fault_B == "b4:0" && (b4_1 = ZERO)
    fault_B == "b4:1" && (b4_1 = ONE)

    _,b5_1 = op3or(null,null,null,s12_1,s13_1,s14_1)
    fault_B == "b5:0" && (b5_1 = ZERO)
    fault_B == "b5:1" && (b5_1 = ONE)

    _,b6_1 = op2and(null,null,s15_1,ni5_1)
    fault_B == "b6:0" && (b6_1 = ZERO)
    fault_B == "b6:1" && (b6_1 = ONE)

    _,b8_1 = op2and(null,null,s23_1,b7_1)
    fault_B == "b8:0" && (b8_1 = ZERO)
    fault_B == "b8:1" && (b8_1 = ONE)

    _,b9_1 = op2or(null,null,s25_1,s26_1)
    fault_B == "b9:0" && (b9_1 = ZERO)
    fault_B == "b9:1" && (b9_1 = ONE)

    _,b10_1 = op3or(null,null,null,s27_1,s28_1,s29_1)
    fault_B == "b10:0" && (b10_1 = ZERO)
    fault_B == "b10:1" && (b10_1 = ONE)

    _,c2_1 = op3and(null,null,null,s7_1,ni5_1,ni3_1)
    fault_B == "c2:0" && (c2_1 = ZERO)
    fault_B == "c2:1" && (c2_1 = ONE)

    _,c3_1 = op2and(null,null,s11_1,i7_1)
    fault_B == "c3:0" && (c3_1 = ZERO)
    fault_B == "c3:1" && (c3_1 = ONE)

    _,c4_1 = op2and(null,null,s15_1,ni5_1)
    fault_B == "c4:0" && (c4_1 = ZERO)
    fault_B == "c4:1" && (c4_1 = ONE)

    _,c6_1 = op2and(null,null,s23_1,ni5_1)
    fault_B == "c6:0" && (c6_1 = ZERO)
    fault_B == "c6:1" && (c6_1 = ONE)

    _,c8_1 = op2and(null,null,s19_1,c5_1)
    fault_B == "c8:0" && (c8_1 = ZERO)
    fault_B == "c8:1" && (c8_1 = ONE)

    _,c14_1 = op2and(null,null,c13_1,s3_1)
    fault_B == "c14:0" && (c14_1 = ZERO)
    fault_B == "c14:1" && (c14_1 = ONE)

    _,c15_1 = op2and(null,null,s27_1,i7_1)
    fault_B == "c15:0" && (c15_1 = ZERO)
    fault_B == "c15:1" && (c15_1 = ONE)

    _,d1_1 = op2and(null,null,s1_1,i2_1)
    fault_B == "d1:0" && (d1_1 = ZERO)
    fault_B == "d1:1" && (d1_1 = ONE)

    _,d2_1 = op3and(null,null,null,s3_1,ni3_1,ni2_1)
    fault_B == "d2:0" && (d2_1 = ZERO)
    fault_B == "d2:1" && (d2_1 = ONE)

    _,d3_1 = op2and(null,null,s5_1,i0_1)
    fault_B == "d3:0" && (d3_1 = ZERO)
    fault_B == "d3:1" && (d3_1 = ONE)

    _,d5_1 = op2and(null,null,s9_1,i2_1)
    fault_B == "d5:0" && (d5_1 = ZERO)
    fault_B == "d5:1" && (d5_1 = ONE)

    _,d6_1 = op2and(null,null,s11_1,ni7_1)
    fault_B == "d6:0" && (d6_1 = ZERO)
    fault_B == "d6:1" && (d6_1 = ONE)

    _,d7_1 = op2and(null,null,s13_1,i0_1)
    fault_B == "d7:0" && (d7_1 = ZERO)
    fault_B == "d7:1" && (d7_1 = ONE)

    _,d9_1 = op2and(null,null,s15_1,ni5_1)
    fault_B == "d9:0" && (d9_1 = ZERO)
    fault_B == "d9:1" && (d9_1 = ONE)

    _,d10_1 = op2and(null,null,s17_1,i2_1)
    fault_B == "d10:0" && (d10_1 = ZERO)
    fault_B == "d10:1" && (d10_1 = ONE)

    _,d11_1 = op2and(null,null,s19_1,ni7_1)
    fault_B == "d11:0" && (d11_1 = ZERO)
    fault_B == "d11:1" && (d11_1 = ONE)

    _,d12_1 = op2and(null,null,s21_1,i0_1)
    fault_B == "d12:0" && (d12_1 = ZERO)
    fault_B == "d12:1" && (d12_1 = ONE)

    _,d13_1 = op3and(null,null,null,s23_1,ni5_1,ni1_1)
    fault_B == "d13:0" && (d13_1 = ZERO)
    fault_B == "d13:1" && (d13_1 = ONE)

    _,d14_1 = op2and(null,null,s25_1,i2_1)
    fault_B == "d14:0" && (d14_1 = ZERO)
    fault_B == "d14:1" && (d14_1 = ONE)

    _,d15_1 = op2and(null,null,s29_1,i0_1)
    fault_B == "d15:0" && (d15_1 = ZERO)
    fault_B == "d15:1" && (d15_1 = ONE)

    _,d27_1 = op2and(null,null,s27_1,ni7_1)
    fault_B == "d27:0" && (d27_1 = ZERO)
    fault_B == "d27:1" && (d27_1 = ONE)

    _,e1_1 = op2and(null,null,s0_1,i1_1)
    fault_B == "e1:0" && (e1_1 = ZERO)
    fault_B == "e1:1" && (e1_1 = ONE)

    _,e2_1 = op2and(null,null,s1_1,ni2_1)
    fault_B == "e2:0" && (e2_1 = ZERO)
    fault_B == "e2:1" && (e2_1 = ONE)

    _,e3_1 = op2and(null,null,s2_1,i2_1)
    fault_B == "e3:0" && (e3_1 = ZERO)
    fault_B == "e3:1" && (e3_1 = ONE)

    _,e4_1 = op3and(null,null,null,s3_1,ni3_1,ni2_1)
    fault_B == "e4:0" && (e4_1 = ZERO)
    fault_B == "e4:1" && (e4_1 = ONE)

    _,e6_1 = op2and(null,null,s5_1,ni0_1)
    fault_B == "e6:0" && (e6_1 = ZERO)
    fault_B == "e6:1" && (e6_1 = ONE)

    _,e7_1 = op2and(null,null,s6_1,i7_1)
    fault_B == "e7:0" && (e7_1 = ZERO)
    fault_B == "e7:1" && (e7_1 = ONE)

    _,e8_1 = op2and(null,null,s8_1,i1_1)
    fault_B == "e8:0" && (e8_1 = ZERO)
    fault_B == "e8:1" && (e8_1 = ONE)

    _,e9_1 = op2and(null,null,s9_1,ni2_1)
    fault_B == "e9:0" && (e9_1 = ZERO)
    fault_B == "e9:1" && (e9_1 = ONE)

    _,e11_1 = op2and(null,null,s11_1,ni7_1)
    fault_B == "e11:0" && (e11_1 = ZERO)
    fault_B == "e11:1" && (e11_1 = ONE)

    _,e13_1 = op2and(null,null,s13_1,ni0_1)
    fault_B == "e13:0" && (e13_1 = ZERO)
    fault_B == "e13:1" && (e13_1 = ONE)

    _,e14_1 = op2and(null,null,s14_1,i7_1)
    fault_B == "e14:0" && (e14_1 = ZERO)
    fault_B == "e14:1" && (e14_1 = ONE)

    _,e15_1 = op2and(null,null,s15_1,ni5_1)
    fault_B == "e15:0" && (e15_1 = ZERO)
    fault_B == "e15:1" && (e15_1 = ONE)

    _,e16_1 = op2and(null,null,s16_1,i1_1)
    fault_B == "e16:0" && (e16_1 = ZERO)
    fault_B == "e16:1" && (e16_1 = ONE)

    _,e17_1 = op2and(null,null,s17_1,ni2_1)
    fault_B == "e17:0" && (e17_1 = ZERO)
    fault_B == "e17:1" && (e17_1 = ONE)

    _,e19_1 = op2and(null,null,s19_1,ni7_1)
    fault_B == "e19:0" && (e19_1 = ZERO)
    fault_B == "e19:1" && (e19_1 = ONE)

    _,e20_1 = op2and(null,null,s20_1,e12_1)
    fault_B == "e20:0" && (e20_1 = ZERO)
    fault_B == "e20:1" && (e20_1 = ONE)

    _,e21_1 = op2and(null,null,s21_1,ni0_1)
    fault_B == "e21:0" && (e21_1 = ZERO)
    fault_B == "e21:1" && (e21_1 = ONE)

    _,e22_1 = op2and(null,null,s22_1,i7_1)
    fault_B == "e22:0" && (e22_1 = ZERO)
    fault_B == "e22:1" && (e22_1 = ONE)

    _,e23_1 = op2and(null,null,s23_1,ni5_1)
    fault_B == "e23:0" && (e23_1 = ZERO)
    fault_B == "e23:1" && (e23_1 = ONE)

    _,e25_1 = op2and(null,null,s25_1,ni2_1)
    fault_B == "e25:0" && (e25_1 = ZERO)
    fault_B == "e25:1" && (e25_1 = ONE)

    _,e26_1 = op2and(null,null,s26_1,i2_1)
    fault_B == "e26:0" && (e26_1 = ZERO)
    fault_B == "e26:1" && (e26_1 = ONE)

    _,e27_1 = op2and(null,null,s27_1,ni7_1)
    fault_B == "e27:0" && (e27_1 = ZERO)
    fault_B == "e27:1" && (e27_1 = ONE)

    _,e28_1 = op2and(null,null,s28_1,e12_1)
    fault_B == "ls7:0" && (ls7_1 = ZERO)
    fault_B == "ls7:1" && (ls7_1 = ONE)

    _,e29_1 = op2and(null,null,s29_1,ni0_1)
    fault_B == "e29:0" && (e29_1 = ZERO)
    fault_B == "e29:1" && (e29_1 = ONE)

    _,e30_1 = op2and(null,null,s30_1,i7_1)
    fault_B == "e30:0" && (e30_1 = ZERO)
    fault_B == "e30:1" && (e30_1 = ONE)

    _,e31_1 = op2and(null,null,s4_1,e5_1)
    fault_B == "e31:0" && (e31_1 = ZERO)
    fault_B == "e31:1" && (e31_1 = ONE)

    _,e32_1 = op2and(null,null,s10_1,e10_1)
    fault_B == "e32:0" && (e32_1 = ZERO)
    fault_B == "e32:1" && (e32_1 = ONE)

    _,e33_1 = op2and(null,null,s12_1,e12_1)
    fault_B == "e33:0" && (e33_1 = ZERO)
    fault_B == "e33:1" && (e33_1 = ONE)

    _,e34_1 = op2and(null,null,s18_1,e18_1)
    fault_B == "e34:0" && (e34_1 = ZERO)
    fault_B == "e34:1" && (e34_1 = ONE)

    _,e35_1 = op2and(null,null,s24_1,e24_1)
    fault_B == "e35:0" && (e35_1 = ZERO)
    fault_B == "e35:1" && (e35_1 = ONE)

    _,f1_1 = op2and(null,null,s12_1,i5_1)
    fault_B == "f1:0" && (f1_1 = ZERO)
    fault_B == "f1:1" && (f1_1 = ONE)

    _,f2_1 = op2and(null,null,s27_1,i4_1)
    fault_B == "f2:0" && (f2_1 = ZERO)
    fault_B == "f2:1" && (f2_1 = ONE)

    _,f3_1 = op2and(null,null,s15_1,i0_1)
    fault_B == "f3:0" && (f3_1 = ZERO)
    fault_B == "f3:1" && (f3_1 = ONE)

    _,f4_1 = op2and(null,null,s27_1,i2_1)
    fault_B == "f4:0" && (f4_1 = ZERO)
    fault_B == "f4:1" && (f4_1 = ONE)

    _,f5_1 = op2and(null,null,s0_1,i7_1)
    fault_B == "f5:0" && (f5_1 = ZERO)
    fault_B == "f5:1" && (f5_1 = ONE)

    _,f6_1 = op2and(null,null,s27_1,i1_1)
    fault_B == "f6:0" && (f6_1 = ZERO)
    fault_B == "f6:1" && (f6_1 = ONE)
    
    # level 5

    _,a8_1 = op2or(null,null,a7_1,s30_1)
    fault_B == "a8:0" && (a8_1 = ZERO)
    fault_B == "a8:1" && (a8_1 = ONE)

    _,a10_1 = op3or(null,null,null,a1_1,a2_1,s16_1)
    fault_B == "a10:0" && (a10_1 = ZERO)
    fault_B == "a10:1" && (a10_1 = ONE)

    _,b11_1 = op2or(null,null,b10_1,s30_1)
    fault_B == "b11:0" && (b11_1 = ZERO)
    fault_B == "b11:1" && (b11_1 = ONE)

    _,b12_1 = op3or(null,null,null,b1_1,b3_1,s8_1)
    fault_B == "b12:0" && (b12_1 = ZERO)
    fault_B == "b12:1" && (b12_1 = ONE)

    _,b13_1 = op3or(null,null,null,s9_1,b4_1,s11_1)
    fault_B == "b13:0" && (b13_1 = ZERO)
    fault_B == "b13:1" && (b13_1 = ONE)

    _,c1_1 = op3or(null,null,null,c14_1,s4_1,s5_1)
    fault_B == "c1:0" && (c1_1 = ZERO)
    fault_B == "c1:1" && (c1_1 = ONE)

    _,c7_1 = op2and(null,null,c2_1,ni2_1)
    fault_B == "c7:0" && (c7_1 = ZERO)
    fault_B == "c7:1" && (c7_1 = ONE)

    _,c16_1 = op3or(null,null,null,c15_1,s28_1,s29_1)
    fault_B == "c16:0" && (c16_1 = ZERO)
    fault_B == "c16:1" && (c16_1 = ONE)

    _,d17_1 = op3or(null,null,null,d1_1,d2_1,d3_1)
    fault_B == "d17:0" && (d17_1 = ZERO)
    fault_B == "d17:1" && (d17_1 = ONE)

    _,d19_1 = op3or(null,null,null,s10_1,d6_1,d7_1)
    fault_B == "d19:0" && (d19_1 = ZERO)
    fault_B == "d19:1" && (d19_1 = ONE)

    _,d20_1 = op3or(null,null,null,s14_1,d9_1,d10_1)
    fault_B == "d20:0" && (d20_1 = ZERO)
    fault_B == "d20:1" && (d20_1 = ONE)

    _,d21_1 = op3or(null,null,null,s18_1,d11_1,d12_1)
    fault_B == "d21:0" && (d21_1 = ZERO)
    fault_B == "d21:1" && (d21_1 = ONE)

    _,d22_1 = op3or(null,null,null,s22_1,d13_1,d14_1)
    fault_B == "d22:0" && (d22_1 = ZERO)
    fault_B == "d22:1" && (d22_1 = ONE)

    _,d23_1 = op3or(null,null,null,s26_1,d15_1,s30_1)
    fault_B == "d23:0" && (d23_1 = ZERO)
    fault_B == "d23:1" && (d23_1 = ONE)

    _,e36_1 = op3or(null,null,null,e1_1,e2_1,e3_1)
    fault_B == "e36:0" && (e36_1 = ZERO)
    fault_B == "e36:1" && (e36_1 = ONE)

    _,e37_1 = op3or(null,null,null,e4_1,e31_1,e6_1)
    fault_B == "e37:0" && (e37_1 = ZERO)
    fault_B == "e37:1" && (e37_1 = ONE)

    _,e39_1 = op2or(null,null,e9_1,e32_1)
    fault_B == "e39:0" && (e39_1 = ZERO)
    fault_B == "e39:1" && (e39_1 = ONE)

    _,e40_1 = op3or(null,null,null,e11_1,e33_1,e13_1)
    fault_B == "e40:0" && (e40_1 = ZERO)
    fault_B == "e40:1" && (e40_1 = ONE)

    _,e41_1 = op3or(null,null,null,e14_1,e15_1,e16_1)
    fault_B == "e41:0" && (e41_1 = ZERO)
    fault_B == "e41:1" && (e41_1 = ONE)

    _,e42_1 = op3or(null,null,null,e17_1,e34_1,e19_1)
    fault_B == "e42:0" && (e42_1 = ZERO)
    fault_B == "e42:1" && (e42_1 = ONE)

    _,e43_1 = op3or(null,null,null,e20_1,e30_1,a9_1)
    fault_B == "e43:0" && (e43_1 = ZERO)
    fault_B == "e43:1" && (e43_1 = ONE)

    _,e44_1 = op3or(null,null,null,e21_1,e22_1,e23_1)
    fault_B == "e44:0" && (e44_1 = ZERO)
    fault_B == "e44:1" && (e44_1 = ONE)

    _,e45_1 = op3or(null,null,null,e35_1,e25_1,e26_1)
    fault_B == "e45:0" && (e45_1 = ZERO)
    fault_B == "e45:1" && (e45_1 = ONE)

    _,e46_1 = op3or(null,null,null,e27_1,e28_1,e29_1)
    fault_B == "e46:0" && (e46_1 = ZERO)
    fault_B == "e46:1" && (e46_1 = ONE)

    _,out4_1 = op2or(null,null,f1_1,f2_1)
    fault_B == "out4:0" && (out4_1 = ZERO)
    fault_B == "out4:1" && (out4_1 = ONE)

    _,out2_1 = op2or(null,null,f3_1,f4_1)
    fault_B == "out2:0" && (out2_1 = ZERO)
    fault_B == "out2:1" && (out2_1 = ONE)

    _,out1_1 = op2or(null,null,f5_1,f6_1)
    fault_B == "out1:0" && (out1_1 = ZERO)
    fault_B == "out1:1" && (out1_1 = ONE)
    
    # level 6

    _,a11_1 = op3or(null,null,null,a10_1,s17_1,a3_1)
    fault_B == "a11:0" && (a11_1 = ZERO)
    fault_B == "a11:1" && (a11_1 = ONE)

    _,b14_1 = op3or(null,null,null,b12_1,b13_1,b5_1)
    fault_B == "b14:0" && (b14_1 = ZERO)
    fault_B == "b14:1" && (b14_1 = ONE)

    _,c9_1 = op3or(null,null,null,c1_1,s6_1,c7_1)
    fault_B == "c9:0" && (c9_1 = ZERO)
    fault_B == "c9:1" && (c9_1 = ONE)

    _,c17_1 = op2or(null,null,c16_1,s30_1)
    fault_B == "c17:0" && (c17_1 = ZERO)
    fault_B == "c17:1" && (c17_1 = ONE)

    _,d18_1 = op3or(null,null,null,s6_1,c7_1,d5_1)
    fault_B == "d18:0" && (d18_1 = ZERO)
    fault_B == "d18:1" && (d18_1 = ONE)

    _,d28_1 = op2or(null,null,d23_1,d27_1)
    fault_B == "d28:0" && (d28_1 = ZERO)
    fault_B == "d28:1" && (d28_1 = ONE)

    _,e38_1 = op3or(null,null,null,e7_1,c7_1,e8_1)
    fault_B == "e38:0" && (e38_1 = ZERO)
    fault_B == "e38:1" && (e38_1 = ONE)

    _,e47_1 = op3or(null,null,null,e36_1,e37_1,e44_1)
    fault_B == "e47:0" && (e47_1 = ZERO)
    fault_B == "e47:1" && (e47_1 = ONE)

    _,e49_1 = op3or(null,null,null,e39_1,e40_1,e41_1)
    fault_B == "e49:0" && (e49_1 = ZERO)
    fault_B == "e49:1" && (e49_1 = ONE)
    
    # level 7

    _,a12_1 = op3or(null,null,null,a11_1,s19_1,a4_1)
    fault_B == "a12:0" && (a12_1 = ZERO)
    fault_B == "a12:1" && (a12_1 = ONE)

    _,b15_1 = op3or(null,null,null,b14_1,b6_1,b8_1)
    fault_B == "b15:0" && (b15_1 = ZERO)
    fault_B == "b15:1" && (b15_1 = ONE)

    _,c10_1 = op3or(null,null,null,c9_1,c3_1,b5_1)
    fault_B == "c10:0" && (c10_1 = ZERO)
    fault_B == "c10:1" && (c10_1 = ONE)

    _,d24_1 = op3or(null,null,null,d17_1,d18_1,d19_1)
    fault_B == "d24:0" && (d24_1 = ZERO)
    fault_B == "d24:1" && (d24_1 = ONE)

    _,e48_1 = op3or(null,null,null,e45_1,e46_1,e38_1)
    fault_B == "e48:0" && (e48_1 = ZERO)
    fault_B == "e48:1" && (e48_1 = ONE)
    
    # level 8

    _,a13_1 = op3or(null,null,null,a12_1,s23_1,a5_1)
    fault_B == "a13:0" && (a13_1 = ZERO)
    fault_B == "a13:1" && (a13_1 = ONE)

    _,b16_1 = op3or(null,null,null,b15_1,s24_1,b9_1)
    fault_B == "b16:0" && (b16_1 = ZERO)
    fault_B == "b16:1" && (b16_1 = ONE)

    _,c11_1 = op3or(null,null,null,c10_1,c4_1,c8_1)
    fault_B == "c11:0" && (c11_1 = ZERO)
    fault_B == "c11:1" && (c11_1 = ONE)

    _,d25_1 = op3or(null,null,null,d24_1,d20_1,d21_1)
    fault_B == "d25:0" && (d25_1 = ZERO)
    fault_B == "d25:1" && (d25_1 = ONE)

    _,e50_1 = op3or(null,null,null,e47_1,e48_1,e49_1)
    fault_B == "e50:0" && (e50_1 = ZERO)
    fault_B == "e50:1" && (e50_1 = ONE)
    
    # level 9

    _,ns1_1 = op3or(null,null,null,e50_1,e42_1,e43_1)
    fault_B == "ns1:0" && (ns1_1 = ZERO)
    fault_B == "ns1:1" && (ns1_1 = ONE)

    _,ns8_1 = op3or(null,null,null,b16_1,b11_1,a9_1)
    fault_B == "ns8:0" && (ns8_1 = ZERO)
    fault_B == "ns8:1" && (ns8_1 = ONE)

    _,a14_1 = op3or(null,null,null,a13_1,a6_1,a8_1)
    fault_B == "a14:0" && (a14_1 = ZERO)
    fault_B == "a14:1" && (a14_1 = ONE)

    _,c12_1 = op3or(null,null,null,c11_1,a4_1,c6_1)
    fault_B == "c12:0" && (c12_1 = ZERO)
    fault_B == "c12:1" && (c12_1 = ONE)

    _,d26_1 = op3or(null,null,null,d25_1,d22_1,d28_1)
    fault_B == "d26:0" && (d26_1 = ZERO)
    fault_B == "d26:1" && (d26_1 = ONE)
    
    # level 10

    _,ns2_1 = op3or(null,null,null,d26_1,s2_1,a9_1)
    fault_B == "ns2:0" && (ns2_1 = ZERO)
    fault_B == "ns2:1" && (ns2_1 = ONE)

    _,ns4_1 = op3or(null,null,null,c12_1,c17_1,a9_1)
    fault_B == "ns4:0" && (ns4_1 = ZERO)
    fault_B == "ns4:1" && (ns4_1 = ONE)

    _,ns16_1 = op2or(null,null,a14_1,a9_1)
    fault_B == "ns16:0" && (ns16_1 = ZERO)
    fault_B == "ns16:1" && (ns16_1 = ONE)

    (out4_1,out2_1,out1_1,ns16_1,ns8_1,ns4_1,ns2_1,ns1_1)

end # of ON_set()


# poke =================================================================
# ======================================================================

# call poke(s, i, fault)

function poke(s::Int64, i::Int64, fault_B::String)
    
    (a,b,c,d,e,f,g,h) = si2ONset(s, i)
    
    (out4_1,out2_1,out1_1,ns16_1,ns8_1,ns4_1,ns2_1,ns1_1) =
        ON_set(a,b,c,d,e,f,g,h, fault_B)
    
    ns = ONset2ns(ns16_1,ns8_1,ns4_1,ns2_1,ns1_1)
#=
    println("fault: ", fault_B)
    conditional_print("out4_1: ", out4_1)
    conditional_print("out2_1: ", out2_1)
    conditional_print("out1_1: ", out1_1)
    println("ns: ", ns)
=#
    return (s, i, ns, out4_1, out2_1, out1_1)
end


function si2ONset(s::Int64, i::Int64)
    
    a::Ptr{Nothing} = ZERO
    b::Ptr{Nothing} = ONE
    c = ()
    
    s == 0 && (c = (a,a,a,a,a))
    s == 1 && (c = (a,a,a,a,b))
    s == 2 && (c = (a,a,a,b,a))
    s == 3 && (c = (a,a,a,b,b))
    s == 4 && (c = (a,a,b,a,a))
    s == 5 && (c = (a,a,b,a,b))
    s == 6 && (c = (a,a,b,b,a))
    s == 7 && (c = (a,a,b,b,b))
    s == 8 && (c = (a,b,a,a,a))
    s == 9 && (c = (a,b,a,a,b))
    s == 10 && (c = (a,b,a,b,a))
    s == 11 && (c = (a,b,a,b,b))
    s == 12 && (c = (a,b,b,a,a))
    s == 13 && (c = (a,b,b,a,b))
    s == 14 && (c = (a,b,b,b,a))
    s == 15 && (c = (a,b,b,b,b))
    s == 16 && (c = (b,a,a,a,a))
    s == 17 && (c = (b,a,a,a,b))
    s == 18 && (c = (b,a,a,b,a))
    s == 19 && (c = (b,a,a,b,b))
    s == 20 && (c = (b,a,b,a,a))
    s == 21 && (c = (b,a,b,a,b))
    s == 22 && (c = (b,a,b,b,a))
    s == 23 && (c = (b,a,b,b,b))
    s == 24 && (c = (b,b,a,a,a))
    s == 25 && (c = (b,b,a,a,b))
    s == 26 && (c = (b,b,a,b,a))
    s == 27 && (c = (b,b,a,b,b))
    s == 28 && (c = (b,b,b,a,a))
    s == 29 && (c = (b,b,b,a,b))
    s == 30 && (c = (b,b,b,b,a))
    s == 31 && (c = (b,b,b,b,b))
    
    i == 0 && (d = (a,a,a))
    i == 1 && (d = (a,a,b))
    i == 2 && (d = (a,b,a))
    i == 3 && (d = (a,b,b))
    i == 4 && (d = (b,a,a))
    i == 5 && (d = (b,a,b))
    i == 6 && (d = (b,b,a))
    i == 7 && (d = (b,b,b))
    
    return tuplejoin(c,d)
end
    

function ONset2ns(ns16_1::Ptr{Nothing}, ns8_1::Ptr{Nothing},
    ns4_1::Ptr{Nothing}, ns2_1::Ptr{Nothing}, ns1_1::Ptr{Nothing})
    
    a::Ptr{Nothing} = ZERO
    b::Ptr{Nothing} = ONE
    ns::Int64 = 0
    
    c = ns16_1
    d = ns8_1
    e = ns4_1
    f = ns2_1
    g = ns1_1
    
    (c,d,e,f,g) == (a,a,a,a,a) && (ns = 0)
    (c,d,e,f,g) == (a,a,a,a,b) && (ns = 1)
    (c,d,e,f,g) == (a,a,a,b,a) && (ns = 2)
    (c,d,e,f,g) == (a,a,a,b,b) && (ns = 3)
    (c,d,e,f,g) == (a,a,b,a,a) && (ns = 4)
    (c,d,e,f,g) == (a,a,b,a,b) && (ns = 5)
    (c,d,e,f,g) == (a,a,b,b,a) && (ns = 6)
    (c,d,e,f,g) == (a,a,b,b,b) && (ns = 7)
    (c,d,e,f,g) == (a,b,a,a,a) && (ns = 8)
    (c,d,e,f,g) == (a,b,a,a,b) && (ns = 9)
    (c,d,e,f,g) == (a,b,a,b,a) && (ns = 10)
    (c,d,e,f,g) == (a,b,a,b,b) && (ns = 11)
    (c,d,e,f,g) == (a,b,b,a,a) && (ns = 12)
    (c,d,e,f,g) == (a,b,b,a,b) && (ns = 13)
    (c,d,e,f,g) == (a,b,b,b,a) && (ns = 14)
    (c,d,e,f,g) == (a,b,b,b,b) && (ns = 15)
    (c,d,e,f,g) == (b,a,a,a,a) && (ns = 16)
    (c,d,e,f,g) == (b,a,a,a,b) && (ns = 17)
    (c,d,e,f,g) == (b,a,a,b,a) && (ns = 18)
    (c,d,e,f,g) == (b,a,a,b,b) && (ns = 19)
    (c,d,e,f,g) == (b,a,b,a,a) && (ns = 20)
    (c,d,e,f,g) == (b,a,b,a,b) && (ns = 21)
    (c,d,e,f,g) == (b,a,b,b,a) && (ns = 22)
    (c,d,e,f,g) == (b,a,b,b,b) && (ns = 23)
    (c,d,e,f,g) == (b,b,a,a,a) && (ns = 24)
    (c,d,e,f,g) == (b,b,a,a,b) && (ns = 25)
    (c,d,e,f,g) == (b,b,a,b,a) && (ns = 26)
    (c,d,e,f,g) == (b,b,a,b,b) && (ns = 27)
    (c,d,e,f,g) == (b,b,b,a,a) && (ns = 28)
    (c,d,e,f,g) == (b,b,b,a,b) && (ns = 29)
    (c,d,e,f,g) == (b,b,b,b,a) && (ns = 30)
    (c,d,e,f,g) == (b,b,b,b,b) && (ns = 31)
    
    return ns
end


# xxx ==================================================================
# ======================================================================

# call xxx(fault,"same"/"diff",

function xxx(fault_B::String, cmp::String, x::Int64)

    s::Int64 = 0
    
    i::Int64 = 0
    
    internal_fault::String = fault_B
    
    empty_string::String = " "
    
    xxx_dict = Dict()
    
    (c, d) = (s, i)
    
    while true
        
        a = poke(c, d, internal_fault)
        
        e = (a[4] == all_terms ? " ONE" : "    ")
        
        f = (a[5] == all_terms ? " ONE" : "    ")
        
        g = (a[6] == all_terms ? " ONE" : "    ")
        
        b = poke(c, d, empty_string)
        
        h = (cmp == "same") & (a == b)
        
        j = (cmp == "diff") & (a != b)
        
        k = (x == a[3])
        
        l = (x == 32)
        
        (h & k) && merge!(xxx_dict, Dict((c, d) => (a[3], e, f, g)))
        (h & l) && merge!(xxx_dict, Dict((c, d) => (a[3], e, f, g)))
        (j & k) && merge!(xxx_dict, Dict((c, d) => (a[3], e, f, g)))
        (j & l) && merge!(xxx_dict, Dict((c, d) => (a[3], e, f, g)))
        
        if (c, d) == (31, 7)
            break
        else
            (c, d) = increment(c, d)
        end
        
    end
    
    return sort(collect(xxx_dict), by=x->x[1])

end


function increment(c::Int64, d::Int64)
    if d != 7
        d += 1
    elseif c != 31
        c += 1
        d = 0
    end
    
    return (c, d)
end
=#

# conditional_print ====================================================
# ======================================================================

function conditional_print(a, b)
    if b == all_terms
        println(a, "ONE")
    elseif b == no_terms
        println(a, "ZERO")
    else
        println(a, allSAT(b))
    end
end

function goo()
    while true
        print("Enter a fault_A name:   ")
        global fault_A= readline()
        f = []
        print("0,0")
        push!(f, peek(0,0))
        # push!(f, peek(0,1))
        # push!(f, peek(0,2))
        print("1,0")
        push!(f, peek(1,0))
        print("1,1")
        push!(f, peek(1,1))
        print("1,2")
        push!(f, peek(1,2))
        print("1,3")
        push!(f, peek(1,3))
        push!(f, peek(1,4))
        push!(f, peek(1,5))
        push!(f, peek(1,6))
        push!(f, peek(1,7))
        push!(f, peek(1,8))
        push!(f, peek(1,9))
        push!(f, peek(1,10))
        push!(f, peek(1,11))
        push!(f, peek(1,12))
        push!(f, peek(1,13))
        push!(f, peek(1,14))
        push!(f, peek(1,15))
        push!(f, peek(1,16))
        push!(f, peek(1,17))
        push!(f, peek(1,18))
        push!(f, peek(1,19))
        push!(f, peek(1,20))
        push!(f, peek(1,21))
        push!(f, peek(1,22))
        push!(f, peek(1,23))
        push!(f, peek(1,24))
        push!(f, peek(1,25))
        push!(f, peek(1,26))
        push!(f, peek(1,27))
        push!(f, peek(1,28))
        push!(f, peek(1,29))
        push!(f, peek(1,30))
        push!(f, peek(1,31))
        push!(f, peek(2,0))
        print("2,1")
        push!(f, peek(2,1))
        push!(f, peek(2,2))
        push!(f, peek(2,3))
        push!(f, peek(2,4))
        push!(f, peek(2,5))
        push!(f, peek(2,6))
        push!(f, peek(2,7))
        push!(f, peek(2,8))
        push!(f, peek(2,9))
        push!(f, peek(2,10))
        push!(f, peek(2,11))
        push!(f, peek(2,12))
        push!(f, peek(2,13))
        push!(f, peek(2,14))
        push!(f, peek(2,15))
        push!(f, peek(2,16))
        push!(f, peek(2,17))
        push!(f, peek(2,18))
        push!(f, peek(2,19))
        push!(f, peek(2,20))
        push!(f, peek(2,21))
        push!(f, peek(2,22))
        push!(f, peek(2,23))
        push!(f, peek(2,24))
        push!(f, peek(2,25))
        push!(f, peek(2,26))
        push!(f, peek(2,27))
        push!(f, peek(2,28))
        push!(f, peek(2,29))
        push!(f, peek(2,30))
        push!(f, peek(2,31))
        push!(f, peek(3,0))
        push!(f, peek(3,1))
        push!(f, peek(3,2))
        push!(f, peek(3,3))
        push!(f, peek(3,4))
        push!(f, peek(3,5))
        push!(f, peek(3,6))
        push!(f, peek(3,7))
        push!(f, peek(3,8))
        push!(f, peek(3,9))
        push!(f, peek(3,10))
        push!(f, peek(3,11))
        push!(f, peek(3,12))
        push!(f, peek(3,13))
        push!(f, peek(3,14))
        push!(f, peek(3,15))
        push!(f, peek(3,16))
        push!(f, peek(3,17))
        push!(f, peek(3,18))
        push!(f, peek(3,19))
        push!(f, peek(3,20))
        push!(f, peek(3,21))
        push!(f, peek(3,22))
        push!(f, peek(3,23))
        push!(f, peek(3,24))
        push!(f, peek(3,25))
        push!(f, peek(3,26))
        push!(f, peek(3,27))
        push!(f, peek(3,28))
        push!(f, peek(3,29))
        push!(f, peek(3,30))
        push!(f, peek(3,31))
        push!(f, peek(4,0))
        push!(f, peek(4,1))
        push!(f, peek(4,2))
        push!(f, peek(4,3))
        push!(f, peek(4,4))
        push!(f, peek(4,5))
        push!(f, peek(4,6))
        push!(f, peek(4,7))
        push!(f, peek(4,8))
        push!(f, peek(4,9))
        push!(f, peek(4,10))
        push!(f, peek(4,11))
        push!(f, peek(4,12))
        push!(f, peek(4,13))
        push!(f, peek(4,14))
        push!(f, peek(4,15))
        push!(f, peek(4,16))
        push!(f, peek(4,17))
        push!(f, peek(4,18))
        push!(f, peek(4,19))
        push!(f, peek(4,20))
        push!(f, peek(4,21))
        push!(f, peek(4,22))
        push!(f, peek(4,23))
        push!(f, peek(4,24))
        push!(f, peek(4,25))
        push!(f, peek(4,26))
        push!(f, peek(4,27))
        push!(f, peek(4,28))
        push!(f, peek(4,29))
        push!(f, peek(4,30))
        push!(f, peek(4,31))
        push!(f, peek(5,0))
        push!(f, peek(5,1))
        push!(f, peek(5,2))
        push!(f, peek(5,3))
        push!(f, peek(5,4))
        push!(f, peek(5,5))
        push!(f, peek(5,6))
        push!(f, peek(5,7))
        push!(f, peek(5,8))
        push!(f, peek(5,9))
        push!(f, peek(5,10))
        push!(f, peek(5,11))
        push!(f, peek(5,12))
        push!(f, peek(5,13))
        push!(f, peek(5,14))
        push!(f, peek(5,15))
        push!(f, peek(5,16))
        push!(f, peek(5,17))
        push!(f, peek(5,18))
        push!(f, peek(5,19))
        push!(f, peek(5,20))
        push!(f, peek(5,21))
        push!(f, peek(5,22))
        push!(f, peek(5,23))
        push!(f, peek(5,24))
        push!(f, peek(5,25))
        push!(f, peek(5,26))
        push!(f, peek(5,27))
        push!(f, peek(5,28))
        push!(f, peek(5,29))
        push!(f, peek(5,30))
        push!(f, peek(5,31))
        push!(f, peek(6,0))
        push!(f, peek(6,1))
        push!(f, peek(6,2))
        push!(f, peek(6,3))
        push!(f, peek(6,4))
        push!(f, peek(6,5))
        push!(f, peek(6,6))
        push!(f, peek(6,7))
        push!(f, peek(6,8))
        push!(f, peek(6,9))
        push!(f, peek(6,10))
        push!(f, peek(6,11))
        push!(f, peek(6,12))
        push!(f, peek(6,13))
        push!(f, peek(6,14))
        push!(f, peek(6,15))
        push!(f, peek(6,16))
        push!(f, peek(6,17))
        push!(f, peek(6,18))
        push!(f, peek(6,19))
        push!(f, peek(6,20))
        push!(f, peek(6,21))
        push!(f, peek(6,22))
        push!(f, peek(6,23))
        push!(f, peek(6,24))
        push!(f, peek(6,25))
        push!(f, peek(6,26))
        push!(f, peek(6,27))
        push!(f, peek(6,28))
        push!(f, peek(6,29))
        push!(f, peek(6,30))
        push!(f, peek(6,31))
        push!(f, peek(7,0))
        push!(f, peek(7,1))
        push!(f, peek(7,2))
        push!(f, peek(7,3))
        push!(f, peek(7,4))
        push!(f, peek(7,5))
        push!(f, peek(7,6))
        push!(f, peek(7,7))
        push!(f, peek(7,8))
        push!(f, peek(7,9))
        push!(f, peek(7,10))
        push!(f, peek(7,11))
        push!(f, peek(7,12))
        push!(f, peek(7,13))
        push!(f, peek(7,14))
        push!(f, peek(7,15))
        push!(f, peek(7,16))
        push!(f, peek(7,17))
        push!(f, peek(7,18))
        push!(f, peek(7,19))
        push!(f, peek(7,20))
        push!(f, peek(7,21))
        push!(f, peek(7,22))
        push!(f, peek(7,23))
        push!(f, peek(7,24))
        push!(f, peek(7,25))
        push!(f, peek(7,26))
        push!(f, peek(7,27))
        push!(f, peek(7,28))
        push!(f, peek(7,29))
        push!(f, peek(7,30))
        push!(f, peek(7,31))

        #println(f)
            
    end
end


# END END END __________________________________________________________
