using System.Globalization;
using System.Text;

namespace Pehape {
	// ReSharper disable once InconsistentNaming
	public static partial class PHP {
		/// <summary>
		/// The Md5Raw function calculates the md5 hash of a string
		/// </summary>
		/// <param name="input">String to be calculated</param>
		/// <returns>16 character binary format</returns>
		public static byte[] Md5Raw(string input) {
			var inputBytes = Encoding.ASCII.GetBytes(input);
#pragma warning disable CA5351
			var hashBytes = System.Security.Cryptography.MD5.HashData(inputBytes);
#pragma warning restore CA5351
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
			foreach (var t in hashBytes) sb.Append(t.ToString("x2", CultureInfo.InvariantCulture));

			return sb.ToString();
		}
	}
}
