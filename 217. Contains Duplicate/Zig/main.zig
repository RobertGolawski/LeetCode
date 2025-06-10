const std = @import("std");
const Allocator = std.mem.Allocator;

pub const TestCase = struct {
    input: []const i32,
    expected: bool,
};

const test_cases = [_]TestCase{
    .{ .input = &.{ 1, 2, 3, 1 }, .expected = true }, // Basic duplicate
    .{ .input = &.{ 1, 2, 3, 4 }, .expected = false }, // No duplicate
    .{ .input = &.{ 1, 1, 1, 3, 3, 4, 3, 2, 4, 2 }, .expected = true }, // Multiple duplicates
    .{ .input = &.{}, .expected = false }, // Empty slice
    .{ .input = &.{1}, .expected = false }, // Single element
    .{ .input = &.{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }, .expected = false }, // No duplicates, longer array
    .{ .input = &.{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 1 }, .expected = true }, // Duplicate at the end
    .{ .input = &.{ 1, 2, 3, 4, 5, 6, 7, 8, 1, 9 }, .expected = true }, // Duplicate in the middle
    .{ .input = &.{ -1, -2, -3, -1 }, .expected = true }, // Negative numbers, duplicate
    .{ .input = &.{ -1, -2, -3, -4 }, .expected = false }, // Negative numbers, no duplicate
    .{ .input = &.{ 0, 0, 0, 0 }, .expected = true }, // All zeros
    .{ .input = &.{ 1, 2, 1, 2, 1, 2 }, .expected = true }, // Alternating duplicates
    .{ .input = &.{ 1, 2, 3, 1, 2, 3 }, .expected = true }, // Duplicate sequences
    .{ .input = &.{ 100000, 100000, 100000 }, .expected = true }, // Large numbers, duplicate
    .{ .input = &.{ 1, 1, 2, 2, 3, 3, 4, 4, 5, 5 }, .expected = true }, // Sorted with adjacent duplicates
    .{ .input = &.{ 5, 4, 3, 2, 1, 5 }, .expected = true }, // Unsorted with duplicate
    .{ .input = &.{ 1, 2, 3, 4, 5, 1, 2, 3, 4, 5 }, .expected = true }, // multiple sequence duplicates
    .{ .input = &.{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 9 }, .expected = true }, // Duplicate at end longer array
};

fn containsDuplicate(allocator: Allocator, nums: []const i32) !bool {
    var set = std.AutoHashMap(i32, void).init(allocator);
    defer set.deinit();

    for (nums) |num| {
        const res = try set.getOrPut(num);
        if (res.found_existing) {
            return true;
        }
    }

    return false;
}

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    for (test_cases, 0..) |t, i| {
        const ans = try containsDuplicate(allocator, t.input);

        if (ans != t.expected) {
            std.debug.print("Test Case {d} failed:\n", .{i});
            std.debug.print("Input: {any}\n", .{t.input});
            std.debug.print("Expected: {}\n", .{t.expected});
            std.debug.print("Got: {}\n", .{ans});
            return;
        }
    }

    std.debug.print("Congratulations! All test cases passed. \n", .{});
}
