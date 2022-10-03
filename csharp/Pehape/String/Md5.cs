using System.Globalization;
using System.Text;

namespace Pehape {
	public static partial class PHP {
		public static byte[] Md5Raw(string input) {
			using var md5 = System.Security.Cryptography.MD5.Create();
			var inputBytes = Encoding.ASCII.GetBytes(input);
			var hashBytes = md5.ComputeHash(inputBytes);
			return hashBytes;
		}

		public static string Md5(string input) {
			var hashBytes = Md5Raw(input);
			var sb = new StringBuilder();
			foreach (var t in hashBytes) {
				sb.Append(t.ToString("X2", CultureInfo.InvariantCulture).ToLowerInvariant());
			}

			return sb.ToString();
		}
	}
}