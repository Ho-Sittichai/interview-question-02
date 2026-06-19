using AuthApi.Data;
using AuthApi.Models;
using BC = BCrypt.Net.BCrypt;

namespace AuthApi.Services;

public class AuthService
{
    private readonly AppDbContext _db;

    public AuthService(AppDbContext db)
    {
        _db = db;
    }

    public async Task<(bool Success, string Message)> RegisterAsync(RegisterRequest request)
    {
        if (string.IsNullOrWhiteSpace(request.Username) || string.IsNullOrWhiteSpace(request.Password))
            return (false, "Username และ Password ห้ามว่าง");

        bool exists = _db.Users.Any(u => u.Username == request.Username);
        if (exists)
            return (false, "Username นี้มีอยู่แล้ว");

        var user = new User
        {
            Username = request.Username,
            Password = BC.HashPassword(request.Password, 12),
            CreatedAt = DateTime.UtcNow
        };

        _db.Users.Add(user);
        await _db.SaveChangesAsync();

        return (true, "สมัครสมาชิกสำเร็จ");
    }
}
