#!/usr/bin/perl
use strict;
use warnings;

my %ascii_to_devanagari = (
    '0' => '०', '1' => '१', '2' => '२', '3' => '३', '4' => '४',
    '5' => '५', '6' => '६', '7' => '७', '8' => '८', '9' => '९',
);

my $in_bangla_map = 0;
my @lines;

while (<>) {
    my $line = $_;
    
    # Detect start of getBanglaTranslations function
    if ($line =~ /func.*getBanglaTranslations.*\{/) {
        $in_bangla_map = 1;
    }
    
    # Detect end of map (closing brace for return statement)
    if ($in_bangla_map && $line =~ /^\s*\}$/ && $lines[-1] =~ /^\s*}/) {
        $in_bangla_map = 0;
    }
    
    # Convert ASCII to Devanagari only within getBanglaTranslations
    # But skip lines that contain "phones" or other keys that shouldn't be translated
    if ($in_bangla_map && $line =~ /:\s*"/ && $line !~ /Available|June|July|August|September|October|November|December|Obsidian|Porcelain|Black|Gold|White/) {
        # Only convert if value side has ASCII digits
        if ($line =~ /:\s*"[^"]*[0-9]/) {
            foreach my $ascii (keys %ascii_to_devanagari) {
                $line =~ s/$ascii/$ascii_to_devanagari{$ascii}/g;
            }
        }
    }
    
    push @lines, $line;
}

print @lines;
