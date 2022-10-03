using System.Globalization;
using System.Text;

namespace Pehape {
	public static partial class PHP {
		/// <summary>
		/// The Md5Raw function calculates the md5 hash of a string
		/// </summary>
		/// <param name="input">String to be calculated</param>
		/// <returns>16 character binary format</returns>
		public static byte[] Md5Raw(string input) {
			using var md5 = System.Security.Cryptography.MD5.Create();
			var inputBytes = Encoding.ASCII.GetBytes(input);
			var hashBytes = md5.ComputeHash(inputBytes);
			return hashBytes;
		}

		/// <summary>
		/// The Md5 function calculates the md5 hash of a string
		/// </summary>
		/// <param name="input">String to be calculated</param>
		/// <returns>32 character hex number</returns>
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