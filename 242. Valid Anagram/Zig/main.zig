const std = @import("std");

pub const TestCase = struct {
    s: []const u8,
    t: []const u8,
    expected: bool,
};

const test_cases = [_]TestCase{
    .{ .s = "anagram", .t = "nagaram", .expected = true },
    .{ .s = "rat", .t = "car", .expected = false },
    .{ .s = "listen", .t = "silent", .expected = true },
    .{ .s = "fried", .t = "fired", .expected = true },
    .{ .s = "aacc", .t = "ccac", .expected = false },
    .{ .s = "qwerty", .t = "qeywrt", .expected = true },
    .{ .s = "a", .t = "b", .expected = false },
    .{ .s = "ab", .t = "a", .expected = false },
    .{ .s = "a", .t = "ab", .expected = false },
    .{ .s = "", .t = "", .expected = true },
    .{ .s = "a", .t = "a", .expected = true },
    .{ .s = "hello", .t = "world", .expected = false },
    .{ .s = "z", .t = "y", .expected = false },
    .{
        .s = "aaaaaaaaaaaaaaaaaaaaaaaaab",
        .t = "bbbbbbbbbbbbbbbbbbbbbbbbba",
        .expected = false,
    },
};

fn isAnagram(s: []const u8, t: []const u8) bool {
    if (s.len != t.len) {
        return false;
    }

    var a = [_]i32{0} ** 26;
    var b = [_]i32{0} ** 26;

    for (s) |letter| {
        a[letter - 'a'] += 1;
    }

    for (t) |letter| {
        b[letter - 'a'] += 1;
    }

    for (0..a.len) |i| {
        if (a[i] != b[i]) {
            return false;
        }
    }

    return true;
}

pub fn main() !void {
    for (test_cases, 0..) |t, i| {
        const ans = isAnagram(t.s, t.t);

        if (ans != t.expected) {
            std.debug.print("Test Case {d} failed:\n", .{i});
            std.debug.print("Input s: \"{s}\"\n", .{t.s});
            std.debug.print("Input t: \"{s}\"\n", .{t.t});
            std.debug.print("Expected: {}\n", .{t.expected});
            std.debug.print("Got: {}\n", .{ans});
            return;
        }
    }

    std.debug.print("Congratulations! All {d} test cases passed.\n", .{test_cases.len});
}

