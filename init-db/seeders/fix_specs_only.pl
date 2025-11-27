#!/usr/bin/perl
use strict;
use warnings;

my %devanagari_to_ascii = (
    '०' => '0', '१' => '1', '२' => '2', '३' => '3', '४' => '4',
    '५' => '5', '६' => '6', '७' => '7', '८' => '8', '९' => '9',
);

my $in_specs_map = 0;
my $brace_count = 0;
my @lines;

while (<>) {
    my $line = $_;
    
    # Detect start of specs map - must be exactly "specs := map[string]string{"
    if ($line =~ /specs\s*:=\s*map\[string\]string\s*\{/) {
        $in_specs_map = 1;
        $brace_count = 1;
    } elsif ($in_specs_map) {
        # Count braces to track when we exit the map
        $brace_count += ($line =~ tr/{//);
        $brace_count -= ($line =~ tr/}//);
        
        # If brace count reaches 0, we're done with the specs map
        if ($brace_count == 0) {
            $in_specs_map = 0;
        }
    }
    
    # Convert Devanagari to ASCII only within specs map
    if ($in_specs_map) {
        foreach my $dev (keys %devanagari_to_ascii) {
            $line =~ s/$dev/$devanagari_to_ascii{$dev}/g;
        }
    }
    
    push @lines, $line;
}

print @lines;
