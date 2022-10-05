using System;
using System.Collections.Generic;
using System.Globalization;
using System.Linq;

namespace Pehape {
	public static partial class PHP {
		/// <summary>
		/// Reverse a string.
		/// </summary>
		/// <param name="str">The string to be reversed.</param>
		/// <returns>Returns the reversed string.</returns>
		public static string Strrev(string str) {
			if (str is null) throw new ArgumentNullException(nameof(str));

			string[] graphemeClusters = GetGraphemeClusters(str).ToArray();
			Array.Reverse(graphemeClusters);
			return string.Join("", graphemeClusters);
		}

		private static IEnumerable<string> GetGraphemeClusters(string str) {
			/*
		  	 this implementation  support reversing Grapheme Clusters correctly which is taken from
			   https://stackoverflow.com/questions/228038/best-way-to-reverse-a-string/15111719#15111719
			   more info: 
			    - https://unicode.org/reports/tr29/, 
			    - https://en.wikipedia.org/wiki/Graphem
				  - https://docs.rs/unicode-reverse/latest/unicode_reverse/
			*/
			TextElementEnumerator enumerator = StringInfo.GetTextElementEnumerator(str);
			while (enumerator.MoveNext()) {
				yield return (string)enumerator.Current;
			}
		}
	}
}