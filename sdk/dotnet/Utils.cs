﻿// Copyright 2016-2019, Pulumi Corporation.

using System;
using System.Text;

namespace Pulumi.Docker
{
    internal static class Utils
    {
        public static (string, string?) GetImageNameAndTag(string baseImageName)
        {
            // From https://docs.docker.com/engine/reference/commandline/tag
            //
            // "A tag name must be valid ASCII and may contain lowercase and uppercase letters, digits,
            // underscores, periods and dashes. A tag name may not start with a period or a dash and may
            // contain a maximum of 128 characters."
            //
            // So it is safe for us to just look for the colon, and consume whatever follows as the tag
            // for the image.

            var lastColon = baseImageName.LastIndexOf(":");
            var imageName = lastColon < 0 ? baseImageName : baseImageName.Substring(0, lastColon);
            var tag = lastColon < 0 ? null : baseImageName.Substring(lastColon + 1);

            return (imageName, tag);
        }

        /// <summary>
        /// Convert an argument array to an argument string for using with Process.StartInfo.Arguments.
        /// </summary>
        public static string EscapeArguments(params string[] args)
        {
            var argEnumerator = args.GetEnumerator();
            var arguments = new StringBuilder();

            if (!argEnumerator.MoveNext())
            {
                return string.Empty;
            }

            arguments.Append(EscapeArguments((string)argEnumerator.Current!));

            while (argEnumerator.MoveNext())
            {
                arguments.Append(' ');
                arguments.Append(EscapeArguments((string)argEnumerator.Current!));
            }

            return arguments.ToString();
        }

        /// <summary>
        /// Convert an argument array to an argument string for using with Process.StartInfo.Arguments.
        /// </summary>
        private static string EscapeArguments(string argument)
        {
            using var characterEnumerator = argument.GetEnumerator();
            var escapedArgument = new StringBuilder();
            var backslashCount = 0;
            var needsQuotes = false;

            while (characterEnumerator.MoveNext())
            {
                switch (characterEnumerator.Current)
                {
                    case '\\':
                        // Backslashes are simply passed through, except when they need
                        // to be escaped when followed by a \", e.g. the argument string
                        // \", which would be encoded to \\\"
                        backslashCount++;
                        escapedArgument.Append('\\');
                        break;

                    case '\"':
                        // Escape any preceding backslashes
                        for (var c = 0; c < backslashCount; c++)
                        {
                            escapedArgument.Append('\\');
                        }

                        // Append an escaped double quote.
                        escapedArgument.Append("\\\"");

                        // Reset the backslash counter.
                        backslashCount = 0;
                        break;

                    case ' ':
                    case '\t':
                        // White spaces are escaped by surrounding the entire string with
                        // double quotes, which should be done at the end to prevent 
                        // multiple wrappings.
                        needsQuotes = true;

                        // Append the whitespace
                        escapedArgument.Append(characterEnumerator.Current);

                        // Reset the backslash counter.
                        backslashCount = 0;
                        break;

                    default:
                        // Reset the backslash counter.
                        backslashCount = 0;

                        // Append the current character
                        escapedArgument.Append(characterEnumerator.Current);
                        break;
                }
            }

            // No need to wrap in quotes
            if (!needsQuotes)
            {
                return escapedArgument.ToString();
            }

            // Prepend the "
            escapedArgument.Insert(0, '"');

            // Escape any preceding backslashes before appending the "
            for (var c = 0; c < backslashCount; c++)
            {
                escapedArgument.Append('\\');
            }

            // Append the final "
            escapedArgument.Append('\"');

            return escapedArgument.ToString();
        }
    }
}
